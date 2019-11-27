import time
from datetime import datetime

from . import status
from .celery import app
from .db import mongo, redis
from .algorithm import compute_battery_statistic, \
    compute_charging_process, compute_working_condition, \
    compute_correlation

TASK_NAME = 'miningTask'
SIG_LIST_NAME = f'{TASK_NAME}:sigList'
WORKING_ID_SET_NAME = f'{TASK_NAME}:workingIdSet'

DATETIME_FORMAT = '%Y-%m-%d %H:%M:%S'


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

    start = time.perf_counter()

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    data_collection = mongo[table_name]
    task_collection = mongo['mining_task']

    task_collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': status.TASK_STATUS_PROCESSING,
        }}
    )
    # 发送信号，websocket相关代码会读取到信号，然后发送数据给前端
    status.send_status_change_sig(SIG_LIST_NAME)

    # 处理数据
    if task_name == '充电过程':
        # projection 字符串中字段的顺序不能变，追加新字段，必须放在最后
        projection = {
            '_id': False,
            '状态号': True,
            '总电压': True,
            'SOC': True,
        }
        compute_alg = compute_charging_process
    elif task_name == '工况':
        projection = {
            '_id': False,
            '时间': True,
            '总电压': True,
            '车速': True,
        }
        compute_alg = compute_working_condition
    elif task_name == '电池统计':
        projection = {
            '_id': False,
            '最高温度电池号': True,
            '最低温度电池号': True,
        }
        compute_alg = compute_battery_statistic
    elif task_name == 'pearson相关系数':
        projection = {
            '_id': False,
            '总电压': True,
            '总电流': True,
            '车速': True,
            'SOC': True,
            '单体最高温度': True,
            '单体最低温度': True,
            '单体最高电压': True,
            '单体最低电压': True,
        }
        compute_alg = compute_correlation
    else:
        return

    if date_range == '所有数据':
        mongo_filter = {}
    else:
        start_date, end_date = date_range.split(' - ')
        mongo_filter = {
            '时间': {
                '$gte': datetime.strptime(start_date, DATETIME_FORMAT),
                '$lte': datetime.strptime(end_date, DATETIME_FORMAT),
            }
        }
    rows = data_collection.find(filter=mongo_filter, projection=projection)
    if rows.count() == 0:
        task_collection.update_one(
            {'taskId': task_id},
            {'$set': {
                'taskStatus': status.TASK_STATUS_FAILURE,
                'comment': '无可用数据',
            }}
        )
    else:
        data = compute_alg(rows)
        used_time = round(time.perf_counter() - start, 2)
        task_collection.update_one(
            {'taskId': task_id},
            {'$set': {
                'taskStatus': status.TASK_STATUS_SUCCESS,
                'comment': f'用时 {used_time}s',
                'data': data
            }}
        )
    status.send_status_change_sig(SIG_LIST_NAME)
    # 删除正在工作集合中自己的 id
    redis.srem(WORKING_ID_SET_NAME, task_id)


@app.task(name='task.mining.stop_compute_model', ignore_result=True)
def stop_compute_model(task_id: str) -> None:
    # 取消一个任务，
    # 如果该任务已执行，那么必须设置 `terminate=True` 才能终止它
    # 如果该任务不存在，也不会报错
    compute_model.AsyncResult(task_id).revoke(terminate=True)
