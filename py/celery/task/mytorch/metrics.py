import torch


def beiqi_accuracy(y_hat, y, alpha: float = 0.02):
    return (torch.abs(y_hat - y) < alpha).sum()
