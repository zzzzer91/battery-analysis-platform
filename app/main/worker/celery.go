package worker

import (
	"battery-anlysis-platform/pkg/conf"
	"battery-anlysis-platform/pkg/worker"
	"github.com/gocelery/gocelery"
)

var Celery *gocelery.CeleryClient

func init() {
	cli, err := worker.InitCelery(&conf.App.Main.Celery)
	if err != nil {
		panic(err)
	}
	Celery = cli
}
