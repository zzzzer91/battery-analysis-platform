package model

import (
	"battery-anlysis-platform/pkg/jtime"
)

type YutongBattery struct {
	ID                        int             `json:"-" gorm:"primary_key"`
	Province                  string          `json:"-" gorm:"type:varchar(100)"`
	City                      string          `json:"-" gorm:"type:varchar(100)"`
	Timestamp                 *jtime.JSONTime `json:"timestamp" gorm:"type:datetime;index"`
	Bty_t_vol                 float64         `json:"bty_t_vol" gorm:"type:decimal(10,2)"`
	Bty_t_curr                float64         `json:"bty_t_curr" gorm:"type:decimal(10,2)"`
	BatterySoc                float64         `json:"battery_soc" gorm:"type:decimal(5,2)"`
	S_b_max_t                 int             `json:"s_b_max_t"`
	Max_t_s_b_num             int             `json:"max_t_s_b_num"`
	S_b_min_t                 int             `json:"s_b_min_t"`
	Min_t_s_b_num             int             `json:"min_t_s_b_num"`
	S_b_max_v                 float64         `json:"s_b_max_v" gorm:"type:decimal(10,2)"`
	Max_v_e_core_num          int             `json:"max_v_e_core_num"`
	S_b_min_v                 float64         `json:"s_b_min_v" gorm:"type:decimal(10,2)"`
	Min_v_e_core_num          int             `json:"min_v_e_core_num"`
	P_t_p                     float64         `json:"p_t_p" gorm:"type:decimal(10,2)"`
	R_t_p                     float64         `json:"r_t_p" gorm:"type:decimal(10,2)"`
	TotalMileage              int             `json:"total_mileage"`
	Bty_sys_rated_capacity    int             `json:"bty_sys_rated_capacity"`
	Bty_sys_rated_consumption int             `json:"bty_sys_rated_consumption"`
	Met_spd                   int             `json:"met_spd"`
	Byt_ma_sys_state          int             `json:"byt_ma_sys_state"`
}
