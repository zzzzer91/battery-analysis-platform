package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	// 设置连接池
	// 空闲
	maxIdleConns = 20
	// 打开
	maxOpenConns = 100
	// 超时
	connMaxLifetime = time.Second * 30
)

func InitMysql(uri string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	// 默认打开 log
	db.LogMode(true)
	// 全局禁用表名复数，默认创建表名时会使用复数
	db.SingularTable(true)

	// 设置连接池
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}
