from typing import Callable

import torch


def beiqi_accuracy(alpha: float) -> Callable:
    return lambda y_hat, y: (torch.abs(y_hat - y) < alpha).sum()
