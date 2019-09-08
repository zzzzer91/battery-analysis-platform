package dao

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/pkg/db"
	"github.com/jinzhu/gorm"
)

// DB 数据库链接单例
var MysqlDB *gorm.DB

// Database 在中间件中初始化mysql链接
func init() {
	d, err := db.InitMysql(conf.Params.MysqlUri)
	if err != nil {
		panic(err)
	}
	MysqlDB = d
}
