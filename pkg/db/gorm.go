package db

import (
	"battery-analysis-platform/pkg/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func InitGorm(gormConf *conf.GormConf) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", gormConf.Uri)
	if err != nil {
		return nil, err
	}
	// 是否打开 log
	db.LogMode(gormConf.LogMode)
	// 全局禁用表名复数，默认创建表名时会使用复数
	db.SingularTable(true)

	// 设置连接池
	db.DB().SetMaxIdleConns(gormConf.MaxIdleConns)
	db.DB().SetMaxOpenConns(gormConf.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(gormConf.ConnMaxLifetime) * time.Second)

	return db, nil
}
