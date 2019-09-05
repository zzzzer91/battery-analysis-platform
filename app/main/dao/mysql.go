package dao

import (
	"battery-anlysis-platform/app/main/model"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var MysqlDB *gorm.DB

// Database 在中间件中初始化mysql链接
func InitMySQL(uri string) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}

	// log
	db.LogMode(true)

	// 全局禁用表名复数，默认创建表名时会使用复数
	db.SingularTable(true)

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	// 自动生成表
	// 会添加没有的字段，但不会修改已有的字段
	db.AutoMigrate(&model.User{})

	MysqlDB = db
}
