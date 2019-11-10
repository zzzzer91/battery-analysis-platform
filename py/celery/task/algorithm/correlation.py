from typing import List, Tuple, Dict, Iterable

import pandas as pd


def compute_correlation(rows: Iterable[Dict]) -> List[Tuple]:
    df = pd.DataFrame(list(rows))
    # 字段顺序不能变
    df = df[
        ['总电压', '总电流', '车速', 'SOC',
         '单体最高温度', '单体最低温度', '单体最高电压', '单体最低电压']
    ]
    results = df.corr('pearson').values.tolist()
    del df

    x, y = 0, 0
    data: List[Tuple[int, int, float]] = []
    for row in results:
        x = 0
        for v in row:
            data.append((x, y, round(v, 3)))
            x += 1
        y += 1
    return data
