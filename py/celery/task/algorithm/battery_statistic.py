from typing import Dict, List, Iterator


def compute_battery_statistic(rows: Iterator[Dict]) -> List[Dict]:
    """计算电池统计数据。"""

    battery_statistic = {}
    for row in rows:
        max_t_s_b_num = row['max_t_s_b_num']
        if max_t_s_b_num is not None:
            battery_statistic.setdefault(max_t_s_b_num, [0, 0])[0] += 1
        min_t_s_b_num = row['min_t_s_b_num']
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