import time

import pandas as pd

from . import status
from .celery import app
from .db import mysql, mongo
from .algorithm import compute_battery_statistic, \
    compute_charging_process, compute_working_condition, \
    compute_correlation

TASK_NAME = 'miningTask'
SIG_LIST_NAME = f'{TASK_NAME}:sigList'


def _gen_sql(need_fields: str, table_name: str, date_range: str) -> str:
    """防 sql 注入的工作由创建任务者检查"""

    if date_range == '所有数据':
        sql = f'SELECT {need_fields} FROM {table_name}'
    else:
        start_date, end_date = date_range.split(' - ')
        sql = f'SELECT {need_fields} FROM {table_name} ' \
              f'WHERE timestamp >= "{start_date}" and timestamp <= "{end_date}"'
    return sql


# `ignore_result=True` 该任务不会将结果保存在 redis，提高性能
@app.task(name='task.mining.compute_model', bind=True, ignore_result=True)
def compute_model(self,
                  task_name: str,
                  # 这几个参数传给 SQl 语句
                  table_name: str,
                  date_range: str) -> None:
    """根据 task_name，选择任务交给 celery 执行。

    :param self: Celery 装饰器中添加 `bind=True` 参数。告诉 Celery 发送一个 self 参数到该函数，
                 可以获取一些任务信息，或更新用 `self.update_stat()` 任务状态。
    :param task_name: 任务名，中文。
    :param table_name: 从哪张表查询数据，表名。
    :param date_range: 数据查询起止日期，格式有：["所有数据"] 和 ["起 - 止"]。
    """

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    collection = mongo['mining_task']

    # 从连接池取一个连接
    mysql_conn = mysql.connect()

    start = time.perf_counter()

    collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': status.TASK_STATUS_PROCESSING,
        }}
    )
    # 发送信号，websocket相关代码会读取到信号，然后发送数据给前端
    status.send_status_change_sig(SIG_LIST_NAME)

    # 处理数据
    data = None
    if task_name == '充电过程':
        # needFields 字符串中字段的顺序不能变，追加新字段，必须放在最后
        need_fields = 'bty_t_vol, bty_t_curr, battery_soc, id, byt_ma_sys_state'
        sql = _gen_sql(need_fields, table_name, date_range)
        # sqlalchemy 默认返回 row 是元组，可以用 `dict(row)` 转换成字典
        rows = mysql_conn.execute(sql)
        if rows.rowcount != 0:
            data = compute_charging_process(rows)
    elif task_name == '工况':
        need_fields = 'timestamp, bty_t_curr, met_spd'
        sql = _gen_sql(need_fields, table_name, date_range)
        rows = mysql_conn.execute(sql)
        if rows.rowcount != 0:
            data = compute_working_condition(rows)
    elif task_name == '电池统计':
        need_fields = 'max_t_s_b_num, min_t_s_b_num'
        sql = _gen_sql(need_fields, table_name, date_range)
        rows = mysql_conn.execute(sql)
        if rows.rowcount != 0:
            data = compute_battery_statistic(rows)
    elif task_name == 'pearson相关系数':
        need_fields = 'bty_t_vol,bty_t_curr,met_spd,battery_soc,' \
                      's_b_max_t,s_b_min_t,s_b_max_v,s_b_min_v'
        sql = _gen_sql(need_fields, table_name, date_range)
        rows = pd.read_sql(sql, con=mysql_conn).corr('pearson').values.tolist()
        data = compute_correlation(rows)
    else:
        return

    # 放回连接
    mysql_conn.close()

    if data is None or len(data) == 0:
        collection.update_one(
            {'taskId': task_id},
            {'$set': {
                'taskStatus': status.TASK_STATUS_FAILURE,
                'comment': '无可用数据',
            }}
        )
    else:
        used_time = round(time.perf_counter() - start, 2)
        collection.update_one(
            {'taskId': task_id},
            {'$set': {
                'taskStatus': status.TASK_STATUS_SUCCESS,
                'comment': f'用时 {used_time}s',
                'data': data
            }}
        )
    status.send_status_change_sig(SIG_LIST_NAME)


@app.task(name='task.mining.stop_compute_model', ignore_result=True)
def stop_compute_model(task_id: str) -> None:
    # 取消一个任务，
    # 如果该任务已执行，那么必须设置 `terminate=True` 才能终止它
    # 如果该任务不存在，也不会报错
    compute_model.AsyncResult(task_id).revoke(terminate=True)
