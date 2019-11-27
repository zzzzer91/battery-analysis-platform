package producer

import (
	"battery-analysis-platform/pkg/conf"
	"github.com/go-redis/redis/v7"
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

// 注意用的 redis 库和 celery 中用的不一样
// key 是 redis 中的集合
func CheckTaskLimit(redis *redis.Client, key string, limit int) bool {
	// 返回集合大小
	ret := redis.SCard(key).Val()
	if ret >= int64(limit) {
		return false
	}
	return true
}

func AddWorkingTaskIdToSet(redis *redis.Client, key string, id string) error {
	if err := redis.SAdd(key, id).Err(); err != nil {
		return err
	}
	return nil
}

func DelWorkingTaskIdFromSet(redis *redis.Client, key string, id string) error {
	if err := redis.SRem(key, id).Err(); err != nil {
		return err
	}
	return nil
}
