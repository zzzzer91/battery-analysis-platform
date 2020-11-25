package celery

import (
	"github.com/gocelery/gocelery"
)

type serviceImpl struct {
	cli *gocelery.CeleryClient
}

func (s *serviceImpl) Delay(task string, args ...interface{}) (string, error) {
	asyncResult, err := s.cli.Delay(task, args...)
	if err != nil {
		return "", err
	}
	return asyncResult.TaskID, nil
}
