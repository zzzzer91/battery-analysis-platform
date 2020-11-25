package producer

import (
	"battery-analysis-platform/app/web/conf"
	producer_celery "battery-analysis-platform/app/web/producer/celery"
	"battery-analysis-platform/pkg/celery"
)

var celeryService producer_celery.Service

func GetCeleryService() producer_celery.Service {
	return celeryService
}

func Init() {
	cli, err := celery.InitCelery(&conf.App.Celery)
	if err != nil {
		panic(err)
	}
	celeryService = producer_celery.NewService(cli)
}
