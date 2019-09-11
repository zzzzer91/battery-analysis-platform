from typing import Dict, List, Iterator


def compute_charging_process(rows: Iterator[Dict]) -> List[Dict]:
    """计算充电过程。"""

    lst1 = []
    i = -1
    pre = -1
    for row in rows:
        if row['byt_ma_sys_state'] != pre:
            pre = row['byt_ma_sys_state']
            lst1.append([])
            i += 1
        lst1[i].append(row)

    lst2 = []
    for row in lst1:
        if row[0]['byt_ma_sys_state'] == 6:
            lst2.append(row)

    data = []
    for i, row in enumerate(lst2, 1):
        max_vol = max(row, key=lambda x: x['bty_t_vol'])['bty_t_vol']
        last_vol = row[-1]['bty_t_vol']
        sub_vol = max_vol - last_vol
        init_soc = row[0]['battery_soc']
        last_soc = row[-1]['battery_soc']
        first_id = row[0]['id']
        last_id = row[-1]['id']
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