package conf

import "battery-analysis-platform/pkg/conf"

type appConf struct {
	Gin    conf.GinConf    `yaml:"gin"`
	Gorm   conf.GormConf   `yaml:"gorm"`
	Mongo  conf.MongoConf  `yaml:"mongo"`
	Celery conf.CeleryConf `yaml:"celery"`
}

var App appConf

func init() {
	if err := conf.Load("go-app-main", &App); err != nil {
		panic(err)
	}
}
