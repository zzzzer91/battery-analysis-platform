import time
import random
from typing import Dict

import torch
from torch import optim

from . import status
from .celery import app
from .db import mongo, redis
from .mytorch.nn import build_nn, train_once
from .mytorch.loss import get_loss
from .mytorch.data import mini_batch
from .mytorch.metrics import beiqi_accuracy

TASK_NAME = 'deeplearningTask'
SIG_LIST_NAME = f'{TASK_NAME}:sigList'
TRAINING_HISTORY_NAME = f'{TASK_NAME}:trainingHistory'


@app.task(name='task.deeplearning.train', bind=True, ignore_result=True)
def train(self, dataset: str, hyper_parameter: Dict):
    """
    普通神经网络训练。

    :param self:
    :param dataset: 数据集名
    :param hyper_parameter: 超参数
    """

    start = time.perf_counter()

    # 用 celery 产生的 id 做 mongo 主键
    task_id = self.request.id

    redis_training_history_loss = f'{TRAINING_HISTORY_NAME}:{task_id}:loss'
    redis_training_history_accuracy = f'{TRAINING_HISTORY_NAME}:{task_id}:accuracy'
    redis_training_history_sig_list = f'{TRAINING_HISTORY_NAME}:{task_id}:sigList'

    data_collection = mongo['beiqi_vehicle']
    task_collection = mongo['deeplearning_task']

    # 固定随机数种子，使结果可以复现
    seed = hyper_parameter['seed']
    random.seed(seed)
    torch.manual_seed(seed)
    torch.cuda.manual_seed(seed)

    if dataset == '北汽_LNBSCU3HXJR884327放电':
        l_temp = (list(d.values()) for d in data_collection.find(
            {'状态号': 2},
            projection={
                '_id': False,
                '时间': False,
                'MSODO总里程': False,
                '状态号': False,
                '动力电池可用能量': False,
                '动力电池可用容量': False,
            }
        ))
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
        y_train = torch.tensor(y[:sample_num], dtype=torch.float)
        x_test = torch.tensor(x[sample_num:], dtype=torch.float)
        y_test = torch.tensor(y[sample_num:], dtype=torch.float)
        alpha_scale = 100
        # 百分比转小数
        if hyper_parameter['outputLayerActivation'].lower() == 'sigmoid':
            y_train /= 100
            y_test /= 100
            alpha_scale = 1
        del x
        del y

        input_dim = x_train.size(1)
        out_dim = 1

        accuracy = beiqi_accuracy(0.02 * alpha_scale)
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

    task_collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': status.TASK_STATUS_PROCESSING,
        }}
    )
    status.send_status_change_sig(SIG_LIST_NAME)

    # 训练
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
        # 存入redis
        redis.rpush(redis_training_history_loss, loss_value_per_epoch)
        redis.rpush(redis_training_history_accuracy, accuracy_value_per_epoch)
        # 发送数据更新信号，用于通知 websocket
        status.send_status_change_sig(redis_training_history_sig_list)

    # 模型评估
    model.eval()
    out = model(x_test)
    accuracy = beiqi_accuracy(0.01 * alpha_scale)
    a_1_count = accuracy(out, y_test).item()
    accuracy = beiqi_accuracy(0.02 * alpha_scale)
    a_2_count = accuracy(out, y_test).item() - a_1_count
    accuracy = beiqi_accuracy(0.03 * alpha_scale)
    a_3_count = accuracy(out, y_test).item() - a_2_count - a_1_count
    accuracy = beiqi_accuracy(0.04 * alpha_scale)
    a_4_count = accuracy(out, y_test).item() - a_3_count - a_2_count - a_1_count
    a_other_count = len(y_test) - a_4_count - a_3_count - a_2_count - a_1_count

    used_time = round(time.perf_counter() - start, 2)
    task_collection.update_one(
        {'taskId': task_id},
        {'$set': {
            'taskStatus': status.TASK_STATUS_SUCCESS,
            'comment': f'用时 {used_time}s',
            'trainingHistory': {
                'loss': loss_history,
                'accuracy': accuracy_history
            },
            'evalResult': {
                'a1Count': a_1_count,
                'a2Count': a_2_count,
                'a3Count': a_3_count,
                'a4Count': a_4_count,
                'aOtherCount': a_other_count,
            }
        }}
    )
    status.send_status_change_sig(SIG_LIST_NAME)
    # 删除暂时存入 Redis 的数据
    redis.delete(redis_training_history_loss)
    redis.delete(redis_training_history_accuracy)
    redis.delete(redis_training_history_sig_list)


@app.task(name='task.deeplearning.stop_train', ignore_result=True)
def stop_train(task_id: str) -> None:
    train.AsyncResult(task_id).revoke(terminate=True)
