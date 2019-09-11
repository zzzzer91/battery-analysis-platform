package dao

import (
	"battery-anlysis-platform/pkg/conf"
	"battery-anlysis-platform/pkg/db"
	"github.com/jinzhu/gorm"
)

var MysqlDB *gorm.DB

func init() {
	d, err := db.InitGorm(&conf.App.Main.Gorm)
	if err != nil {
		panic(err)
	}
	MysqlDB = d
}
