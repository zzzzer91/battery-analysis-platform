from typing import List, Iterable


def compute_correlation(rows: Iterable) -> List[List]:
    x, y = 0, 0
    data = []
    for row in rows:
        x = 0
        for v in row:
            data.append([x, y, round(v, 3)])
            x += 1
        y += 1
    return data
