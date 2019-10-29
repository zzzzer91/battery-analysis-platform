from typing import Iterable

from torch.utils import data


def mini_batch(x, y, batch_size: int, shuffle: bool = True) -> Iterable:
    """这个函数只需要调用一次。
    虽然感觉返回的是生成器，但这个生成器耗尽后，会重置。
    """

    # 将训练数据的特征和标签组合
    # input的数据类型必须是float
    train_dataset = data.TensorDataset(x, y)
    # 随机读取小批量
    train_data_iter = data.DataLoader(train_dataset, batch_size, shuffle=shuffle)
    return train_data_iter
