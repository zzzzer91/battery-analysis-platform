package worker

import (
	"github.com/gocelery/gocelery"
)

func InitCelery(uri string) (*gocelery.CeleryClient, error) {
	cli, err := gocelery.NewCeleryClient(
		gocelery.NewRedisCeleryBroker(uri),
		gocelery.NewRedisCeleryBackend(uri),
		1,
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
