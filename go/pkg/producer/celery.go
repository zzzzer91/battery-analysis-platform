package producer

import (
	"battery-analysis-platform/pkg/conf"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func InitCelery(celeryConf *conf.CeleryConf) (*gocelery.CeleryClient, error) {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(celeryConf.RedisUri)
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
