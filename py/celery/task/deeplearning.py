import time
import random
from typing import List, Dict

import torch
from torch import optim

from .celery import app
from .db import mongo
from .mytorch.nn import build_nn, train_once
from .mytorch.loss import get_loss
from .mytorch.data import mini_batch
from .mytorch.metrics import beiqi_accuracy


@app.task(name='task.deeplearning.train', bind=True, ignore_result=True)
def train(self, dataset: str, hyper_parameter: Dict):
    """
    普通神经网络训练。

    :param self:
    :param dataset: 数据集名
    :param hyper_parameter: 超参数
    """

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    data_collection = mongo['beiqi_vehicle']
    task_collection = mongo['deeplearning_task']

    # 固定随机数种子，使结果可以复现
    seed = hyper_parameter['seed']
    random.seed(seed)
    torch.manual_seed(seed)
    torch.cuda.manual_seed(seed)

    start = time.perf_counter()

    if dataset == '北汽_LNBSCU3HXJR884327放电':
        l_temp = [list(d.values()) for d in data_collection.find(
            {'动力电池充放电状态': 2},
            projection={
                '_id': False,
                '时间': False,
                'MSODO总里程': False,
                '动力电池充放电状态': False,
                '动力电池可用能量': False,
                '动力电池可用容量': False,
            }
        )]
        random.shuffle(l_temp)
        x = []
        y = []
        for v in l_temp:
            x.append(v[1:])
            y.append(v[0:1])
        del l_temp

        # 划分训练，测试数据集
        sample_num = 72000
        x_train = torch.tensor(x[:sample_num], dtype=torch.float)
        y_train = torch.tensor(y[:sample_num], dtype=torch.float) / 100  # 百分比转小数
        x_test = torch.tensor(x[sample_num:], dtype=torch.float)
        y_test = torch.tensor(y[sample_num:], dtype=torch.float) / 100
        del x
        del y

        input_dim = x_train.size(1)
        out_dim = 1

        accuracy = beiqi_accuracy
    else:
        raise ValueError('Non-supported dataset')

    train_data_iter = mini_batch(x_train, y_train, hyper_parameter['batchSize'])

    model = build_nn(
        hyper_parameter['hiddenLayerStructure'],
        input_dim, out_dim,
        hyper_parameter['outputLayerActivation']
    )
    criterion = get_loss(hyper_parameter['loss'])
    optimizer = optim.Adam(model.parameters(), lr=hyper_parameter['learningRate'])

    loss_history = []
    accuracy_history = []
    model.train()
    for i in range(1, hyper_parameter['epochs'] + 1):
        loss_value, accuracy_value = train_once(
            model, train_data_iter, optimizer, criterion, accuracy
        )
        loss_value_per_epoch = round(loss_value / sample_num, 4)
        accuracy_value_per_epoch = round(accuracy_value / sample_num, 4)
        loss_history.append(loss_value_per_epoch)
        accuracy_history.append(accuracy_value_per_epoch)

    used_time = round(time.perf_counter() - start, 2)
    task_collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': '完成',
            'comment': f'用时 {used_time}s',
            'trainingHistory': {
                'loss': loss_history,
                'accuracy': accuracy_history
            }
        }}
    )


@app.task(name='task.deeplearning.stop_train', ignore_result=True)
def stop_train(task_id: str) -> None:
    train.AsyncResult(task_id).revoke(terminate=True)
