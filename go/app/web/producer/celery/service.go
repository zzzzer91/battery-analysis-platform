package celery

import "github.com/gocelery/gocelery"

type Service interface {
	Delay(task string, args ...interface{}) (string, error)
}

func NewService(cli *gocelery.CeleryClient) Service {
	return &serviceImpl{cli}
}
