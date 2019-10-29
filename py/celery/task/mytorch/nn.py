from typing import List, Dict, Tuple

from torch import nn

from .activation import get_activation


def build_nn_architecture(hidden_layer_structure: List[Dict],
                           input_dim: int,
                           output_dim: int,
                           output_layer_activation: str) -> List[Dict]:
    """
    example:
    nn_architecture = [
        {"input_dim": 28*28, "output_dim": 64, "activation": "relu"},
        {"input_dim": 64, "output_dim": 64, "activation": "relu"},
        {"input_dim": 64, "output_dim": 10, "activation": "softmax"},
    ]
    """
    nn_structure = []
    pre = input_dim
    for layer in hidden_layer_structure:
        nn_structure.append({
            'input_dim': pre,
            'output_dim': layer['neurons'],
            'activation': layer['activation']
        })
        pre = layer['neurons']
    nn_structure.append({
        'input_dim': pre,
        'output_dim': output_dim,
        'activation': output_layer_activation
    })
    return nn_structure


def build_nn(hidden_layer_structure: List[Dict],
             input_dim: int,
             output_dim: int,
             output_layer_activation: str):
    """
    example:
    hidden_layer_structure = [
        {"neurons": 64, "activation": "relu"},
        {"neurons": 64, "activation": "relu"},
        {"neurons": 10, "activation": "softmax"},
    ]
    """

    nn_structure = build_nn_architecture(
        hidden_layer_structure, input_dim, output_dim, output_layer_activation
    )

    length = len(nn_structure)
    modules = []
    for i, l in enumerate(nn_structure, 1):
        modules.append(nn.Linear(l['input_dim'], l['output_dim']))
        # 最后一层不需要batch norm
        if i != length:
            modules.append(nn.BatchNorm1d(l['output_dim']))
        a = l.get('activation')
        if a is None or a.strip() == '' or a.strip().lower() == 'linear':
            continue
        modules.append(get_activation(a.strip()))
        # 最后一层不要dropout
    #         if i != length:
    #             modules.append(nn.Dropout(0.5))
    return nn.Sequential(*modules)


def train_once(model, data_iter, optimizer, criterion, accuracy) -> Tuple[float, float]:
    """训练一次。"""

    loss_value: float = 0.
    accuracy_value: float = 0.
    for x, y in data_iter:
        # 注意，从 0.4 版本 Variable 就和 Tensor 合并了，所以不再需要
        # Forward pass: Compute predicted y by passing x to the model
        out = model(x)

        # Compute and print loss
        loss = criterion(out, y)
        # item 取出唯一值，转为python的int类型
        loss_value += loss.item()
        accuracy_value += accuracy(out, y).item()

        # Zero gradients, perform a backward pass, and update the weights.
        optimizer.zero_grad()
        loss.backward()
        optimizer.step()
    return loss_value, accuracy_value
