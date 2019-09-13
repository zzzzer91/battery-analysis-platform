package model

//type YutongBattery struct {
//	Id                        int             `json:"-" gorm:"primary_key"`
//	Province                  string          `json:"-" gorm:"type:varchar(100)"`
//	City                      string          `json:"-" gorm:"type:varchar(100)"`
//	Timestamp                 *jtime.JSONTime `json:"timestamp" gorm:"type:datetime;index"`
//	Bty_t_vol                 float64         `json:"bty_t_vol" gorm:"type:decimal(10,2)"`
//	Bty_t_curr                float64         `json:"bty_t_curr" gorm:"type:decimal(10,2)"`
//	BatterySoc                float64         `json:"battery_soc" gorm:"type:decimal(5,2)"`
//	S_b_max_t                 int             `json:"s_b_max_t"`
//	Max_t_s_b_num             int             `json:"max_t_s_b_num"`
//	S_b_min_t                 int             `json:"s_b_min_t"`
//	Min_t_s_b_num             int             `json:"min_t_s_b_num"`
//	S_b_max_v                 float64         `json:"s_b_max_v" gorm:"type:decimal(10,2)"`
//	Max_v_e_core_num          int             `json:"max_v_e_core_num"`
//	S_b_min_v                 float64         `json:"s_b_min_v" gorm:"type:decimal(10,2)"`
//	Min_v_e_core_num          int             `json:"min_v_e_core_num"`
//	P_t_p                     float64         `json:"p_t_p" gorm:"type:decimal(10,2)"`
//	R_t_p                     float64         `json:"r_t_p" gorm:"type:decimal(10,2)"`
//	TotalMileage              int             `json:"total_mileage"`
//	Bty_sys_rated_capacity    int             `json:"bty_sys_rated_capacity"`
//	Bty_sys_rated_consumption int             `json:"bty_sys_rated_consumption"`
//	Met_spd                   int             `json:"met_spd"`
//	Byt_ma_sys_state          int             `json:"byt_ma_sys_state"`
//}

type MysqlTable struct {
	Name        string
	FieldToName map[string]string
}

var BatteryMysqlNameToTable map[string]MysqlTable

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

	BatteryMysqlNameToTable = map[string]MysqlTable{
		"宇通_4F37195C1A908CFBE0532932A8C0EECB": {
			Name: "yutong_vehicle1", FieldToName: yutongFieldToName,
		},
	}
}
