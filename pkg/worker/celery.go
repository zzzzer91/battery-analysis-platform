package worker

import (
	"battery-anlysis-platform/pkg/conf"
	"github.com/gocelery/gocelery"
)

func InitCelery(celeryConf *conf.CeleryConf) (*gocelery.CeleryClient, error) {
	cli, err := gocelery.NewCeleryClient(
		gocelery.NewRedisCeleryBroker(celeryConf.BrokerUri),
		gocelery.NewRedisCeleryBackend(celeryConf.BackendUri),
		1,
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
