package worker

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/pkg/worker"
	"github.com/gocelery/gocelery"
)

var Celery *gocelery.CeleryClient

func init() {
	cli, err := worker.InitCelery(&conf.App.Celery)
	if err != nil {
		panic(err)
	}
	Celery = cli
}
