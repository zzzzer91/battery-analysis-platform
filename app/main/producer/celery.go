package producer

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/pkg/producer"
	"github.com/gocelery/gocelery"
)

var Celery *gocelery.CeleryClient

func init() {
	cli, err := producer.InitCelery(&conf.App.Celery)
	if err != nil {
		panic(err)
	}
	Celery = cli
}
