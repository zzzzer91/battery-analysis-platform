from typing import Dict, List, Tuple, Iterable


def compute_charging_process(rows: Iterable[Dict]) -> List[Dict]:
    """计算充电过程。注意：若 row 字段顺序改变，算法必须修改，所以添加新字段要追加在最后"""

    lst1: List[List[Dict]] = []
    i = -1
    pre = -1  # 状态号
    for row in rows:
        if row['状态号'] != pre:  # 状态号与前一个数据不同，说明冲放电状态改变
            pre = row['状态号']
            lst1.append([])
            i += 1
        lst1[i].append(row)

    lst2 = []
    for row in lst1:
        if row[0]['状态号'] == 6:  # 充电
            lst2.append(row)

    data: List[Dict] = []
    for i, row in enumerate(lst2, 1):
        max_vol = max(row, key=lambda x: x['总电压'])['总电压']
        last_vol = row[-1]['总电压']
        sub_vol = max_vol - last_vol
        init_soc = row[0]['SOC']
        last_soc = row[-1]['SOC']
        data.append({
            '充电序号': i,
            '最大电压': round(float(max_vol), 1),
            '终止电压': round(float(last_vol), 1),
            '压差': round(float(sub_vol), 1),
            '初始SOC': round(float(init_soc), 1),
            '终止SOC': round(float(last_soc), 1),
        })

    return data
