package cache

import (
	"battery-analysis-platform/app/web/cache/redis"
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/pkg/db"
)

var redisService redis.Service

func GetRedisService() redis.Service {
	return redisService
}

func Init() {
	cli, err := db.InitRedis(&conf.App.Redis)
	if err != nil {
		panic(err)
	}
	redisService = redis.NewService(cli)
}
