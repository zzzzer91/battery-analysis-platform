package dao

import (
	"battery-anlysis-platform/pkg/conf"
	"battery-anlysis-platform/pkg/db"
	"github.com/jinzhu/gorm"
)

var MysqlDB *gorm.DB

func init() {
	d, err := db.InitMysql(conf.App.Main.MysqlUri)
	if err != nil {
		panic(err)
	}
	MysqlDB = d
}
