from typing import Dict, List, Iterable


def compute_battery_statistic(rows: Iterable[Dict]) -> List[Dict]:
    """计算电池统计数据。注意：若 row 字段顺序改变，算法必须修改，所以添加新字段要追加在最后"""

    battery_statistic = {}
    for row in rows:
        max_t_s_b_num = row['最高温度电池号']
        if max_t_s_b_num is not None:
            battery_statistic.setdefault(max_t_s_b_num, [0, 0])[0] += 1
        min_t_s_b_num = row['最低温度电池号']
        if min_t_s_b_num is not None:
            battery_statistic.setdefault(min_t_s_b_num, [0, 0])[1] += 1

    battery_statistic_sorted = sorted(battery_statistic.items(), key=lambda x: x[0])
    data = []
    for number, (max_t_count, min_t_count) in battery_statistic_sorted:
        data.append({
            '电池号': f'{number}号',
            '最大温度次数': max_t_count,
            '最小温度次数': min_t_count
        })
    return data
