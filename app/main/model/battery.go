package model

import "time"

type YutongBattery struct {
	ID                        int        `gorm:"primary_key"`
	Province                  string     `gorm:"type:varchar(100)"`
	City                      string     `gorm:"type:varchar(100)"`
	Timestamp                 *time.Time `gorm:"index"`
	Bty_t_vol                 float64    `gorm:"type:decimal(10,2)"`
	Bty_t_curr                float64    `gorm:"type:decimal(10,2)"`
	BatterySoc                float64    `gorm:"type:decimal(5,2)"`
	S_b_max_t                 int
	Max_t_s_b_num             int
	S_b_min_t                 int
	Min_t_s_b_num             int
	S_b_max_v                 float64 `gorm:"type:decimal(10,2)"`
	Max_v_e_core_num          int
	S_b_min_v                 float64 `gorm:"type:decimal(10,2)"`
	Min_v_e_core_num          int
	P_t_p                     float64 `gorm:"type:decimal(10,2)"`
	R_t_p                     float64 `gorm:"type:decimal(10,2)"`
	TotalMileage              int
	Bty_sys_rated_capacity    int
	Bty_sys_rated_consumption int
	Met_spd                   int
	Byt_ma_sys_state          int
}
