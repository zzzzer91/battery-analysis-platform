package db

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/pkg/db"
	"github.com/jinzhu/gorm"
)

var Gorm *gorm.DB

func init() {
	d, err := db.InitGorm(&conf.App.Gorm)
	if err != nil {
		panic(err)
	}
	Gorm = d
}
