from .db import redis


TASK_STATUS_PREPARING = 0
TASK_STATUS_PROCESSING = 1
TASK_STATUS_SUCCESS = 6
TASK_STATUS_FAILURE = 7


_SIG_VALUE = 1


# 通知 websocket，redis 中数据有更新了
def send_status_change_sig(list_name: str):
    # 防止元素累积
    redis.rpop(list_name)
    redis.rpush(list_name, _SIG_VALUE)
