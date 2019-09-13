from typing import Dict, List, Tuple, Iterator


def compute_battery_statistic(rows: Iterator[Tuple]) -> List[Dict]:
    """计算电池统计数据。注意：若 row 字段顺序改变，算法必须修改，所以添加新字段要追加在最后"""

    battery_statistic = {}
    for row in rows:
        max_t_s_b_num = row[0]
        if max_t_s_b_num is not None:
            battery_statistic.setdefault(max_t_s_b_num, [0, 0])[0] += 1
        min_t_s_b_num = row[1]
        if min_t_s_b_num is not None:
            battery_statistic.setdefault(min_t_s_b_num, [0, 0])[1] += 1

    battery_statistic_sorted = sorted(battery_statistic.items(), key=lambda x: x[0])
    data = []
    for number, (max_t_count, min_t_count) in battery_statistic_sorted:
        data.append({
            'number': f'{number}号',
            'max_t_count': max_t_count,
            'min_t_count': min_t_count
        })
    return data
