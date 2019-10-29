"""
create:   2019-10-19
"""

from typing import Callable

from torch import nn


def get_activation(name: str) -> Callable:
    name = name.lower()
    if name == 'relu':
        return nn.ReLU()
    elif name == 'leaky relu':
        return nn.LeakyReLU()
    elif name == 'sigmoid':
        return nn.Sigmoid()
    elif name == 'linear':
        return lambda x: x
    else:
        raise ValueError('Non-supported activation function')
