package mapping

type MysqlTable struct {
	Name        string
	FieldToName map[string]string
}

var MysqlNameToTable map[string]MysqlTable

func init() {
	yutongFieldToName := map[string]string{
		"timestamp":        "时间",
		"bty_t_vol":        "总电压",
		"bty_t_curr":       "总电流",
		"met_spd":          "车速",
		"p_t_p":            "正向累计电量",
		"r_t_p":            "反向累计电量",
		"total_mileage":    "总里程",
		"battery_soc":      "SOC",
		"byt_ma_sys_state": "状态号",
		"s_b_max_t":        "单体最高温度",
		"s_b_min_t":        "单体最低温度",
		"s_b_max_v":        "单体最高电压",
		"s_b_min_v":        "单体最低电压",
		"max_t_s_b_num":    "最高温度电池号",
		"min_t_s_b_num":    "最低温度电池号",
		"max_v_e_core_num": "最高电压电池号",
		"min_v_e_core_num": "最低电压电池号",
	}

	MysqlNameToTable = map[string]MysqlTable{
		"yutong_4F37195C1A908CFBE0532932A8C0EECB": {
			Name: "yutong_vehicle1", FieldToName: yutongFieldToName,
		},
	}
}
