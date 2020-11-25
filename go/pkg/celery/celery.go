package celery

import (
	"battery-analysis-platform/pkg/conf"
	"github.com/gocelery/gocelery"
	redigo "github.com/gomodule/redigo/redis"
)

func InitCelery(celeryConf *conf.CeleryConf) (*gocelery.CeleryClient, error) {
	redisPool := &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.DialURL(celeryConf.RedisUri)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	cli, err := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		gocelery.NewRedisBackend(redisPool),
		1, // client 随便填
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
