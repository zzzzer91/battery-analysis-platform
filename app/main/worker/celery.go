package worker

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/pkg/worker"
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
