package conf

import "battery-anlysis-platform/pkg/conf"

type appConf struct {
	Gin    conf.GinConf    `yaml:"gin"`
	Gorm   conf.GormConf   `yaml:"gorm"`
	Mongo  conf.MongoConf  `yaml:"mongo"`
	Celery conf.CeleryConf `yaml:"celery"`
}

var App appConf

func init() {
	if err := conf.Load("main", &App); err != nil {
		panic(err)
	}
}
