import time
from typing import List, Dict

from .celery import app
from .db import mongo


@app.task(name='task.deeplearning.train', bind=True, ignore_result=True)
def train(self,
          dataset: str,
          loss: str,
          epochs: int,
          batch_size: int,
          nn_architecture: List[Dict]):
    """
    普通神经网络训练。

    :param self:
    :param dataset: 数据集名
    :param loss: 损失函数名
    :param epochs: 迭代次数
    :param batch_size:
    :param nn_architecture: 神经网络结构
    """

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    collection = mongo['mining_task']

    start = time.perf_counter()

    # TODO


@app.task(name='task.deeplearning.stop_train', ignore_result=True)
def stop_train(task_id: str) -> None:
    train.AsyncResult(task_id).revoke(terminate=True)
