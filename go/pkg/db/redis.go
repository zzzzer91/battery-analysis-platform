package db

import (
	"battery-analysis-platform/pkg/conf"
	"github.com/go-redis/redis/v7"
)

func InitRedis(redisConf *conf.RedisConf) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Uri,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
