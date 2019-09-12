from typing import Dict, List, Tuple, Iterator


def compute_charging_process(rows: Iterator[Tuple]) -> List[Dict]:
    """计算充电过程。注意：若 row 字段顺序改变，算法必须修改，所以添加新字段要追加在最后"""

    lst1: List[List[Tuple]] = []
    i = -1
    pre = -1
    for row in rows:
        if row[4] != pre:
            pre = row[4]
            lst1.append([])
            i += 1
        lst1[i].append(row)

    lst2 = []
    for row in lst1:
        if row[0][4] == 6:
            lst2.append(row)

    data = []
    for i, row in enumerate(lst2, 1):
        max_vol = max(row, key=lambda x: x[0])[0]
        last_vol = row[-1][0]
        sub_vol = max_vol - last_vol
        init_soc = row[0][2]
        last_soc = row[-1][2]
        first_id = row[0][3]
        last_id = row[-1][3]
        data.append({
            'index': i,
            'max_vol': float(max_vol),
            'last_vol': float(last_vol),
            'sub_vol': float(sub_vol),
            'init_soc': float(init_soc),
            'last_soc': float(last_soc),
            'first_id': first_id,
            'last_id': last_id,
        })

    return data
