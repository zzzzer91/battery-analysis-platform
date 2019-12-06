package conf

import "battery-analysis-platform/pkg/conf"

type appConf struct {
	Gin    conf.GinConf    `yaml:"gin"`
	Gorm   conf.GormConf   `yaml:"gorm"`
	Mongo  conf.MongoConf  `yaml:"mongo"`
	Celery conf.CeleryConf `yaml:"celery"`
	Redis  conf.RedisConf  `yaml:"redis"`
}

var App *appConf

func init() {
	app := appConf{}
	if err := conf.Load("go-app-main", &app); err != nil {
		panic(err)
	}
	App = &app
}
