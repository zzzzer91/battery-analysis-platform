from typing import Callable

import torch
from torch import nn


def get_loss(name: str, alpha: float = 0.02) -> Callable:
    name = name.lower()
    if name == 'mse':
        return nn.MSELoss()
    elif name == 'l1':
        return nn.L1Loss()
    elif name == 'sickle-l1':
        return lambda x, y: torch.sum(torch.clamp(torch.abs(x - y), min=alpha))
    else:
        raise ValueError('Non-supported loss function')
