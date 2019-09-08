package dao

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/pkg/db"
	"github.com/jinzhu/gorm"
)

var MysqlDB *gorm.DB

func init() {
	d, err := db.InitMysql(conf.Params.MysqlUri)
	if err != nil {
		panic(err)
	}
	MysqlDB = d
}
