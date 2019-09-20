import time

from sqlalchemy import text

from .celery import app
from .db import mysql, mongo
from .algorithm import compute_battery_statistic
from .algorithm import compute_charging_process
from .algorithm import compute_working_condition


# `ignore_result=True` 该任务不会将结果保存在 redis，提高性能
@app.task(name='task.compute_model', bind=True, ignore_result=True)
def compute_model(self,
                  task_name: str,
                  # 这三个参数传给 SQl 语句
                  need_fields: str,
                  table_name: str,
                  request_params: str) -> None:
    """根据 task_name，选择任务交给 celery 执行。

    :param self: Celery 装饰器中添加 `bind=True` 参数。告诉 Celery 发送一个 self 参数到该函数，
                 可以获取一些任务信息，或更新用 `self.update_stat()` 任务状态。
    :param need_fields: 计算需要的字段。
    :param task_name: 任务名，中文。
    :param table_name: 从哪张表查询数据，表名。
    :param request_params: 数据查询起止日期，格式有：["所有数据"] 和 ["起 - 止"]。
    """

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    if task_name == '充电过程':
        compute_alg = compute_charging_process
    elif task_name == '工况':
        compute_alg = compute_working_condition
    elif task_name == '电池统计':
        compute_alg = compute_battery_statistic
    else:
        return

    start = time.perf_counter()

    # 从连接池取一个连接
    mysql_conn = mysql.connect()
    if request_params == '所有数据':
        rows = mysql_conn.execute(
            'SELECT '
            f'{need_fields} '
            f'FROM {table_name}'
        )
    else:
        start_date, end_date = request_params.split(' - ')
        rows = mysql_conn.execute(
            text(
                'SELECT '
                f'{need_fields} '
                f'FROM {table_name} '
                'WHERE timestamp >= :start_date and timestamp <= :end_date'
            ),
            {'start_date': start_date, 'end_date': end_date}
        )

    collection = mongo['mining_tasks']
    if rows.rowcount == 0:
        collection.update_one(
            {'taskId': task_id},
            {'$set': {
                'taskStatus': '失败',
                'comment': '无可用数据',
            }}
        )
        return

    # 处理数据
    # sqlalchemy 默认返回 row 是元组，可以用 `dict(row)` 转换成字典
    data = compute_alg(rows)
    # 放回连接
    mysql_conn.close()

    used_time = round(time.perf_counter() - start, 2)

    collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': '完成',
            'comment': f'用时 {used_time}s',
            'data': data
        }}
    )


@app.task(name='task.stop_compute_model', ignore_result=True)
def stop_compute_model(task_id: str) -> None:
    # 取消一个任务，
    # 如果该任务已执行，那么必须设置 `terminate=True` 才能终止它
    # 如果该任务不存在，也不会报错
    compute_model.AsyncResult(task_id).revoke(terminate=True)
