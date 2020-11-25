package redis

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/model"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

type serviceImpl struct {
	cli *redis.Client
}

func (s *serviceImpl) Del(strs ...string) error {
	return s.cli.Del(strs...).Err()
}

func (s *serviceImpl) LRange(key string, start, stop int64) ([]string, error) {
	return s.cli.LRange(key, start, stop).Result()
}

func (s *serviceImpl) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return s.cli.BLPop(timeout, keys...).Result()
}

func (s *serviceImpl) AddUserToCache(user *model.User) error {
	// 存储 JSON 序列化的数据
	jd, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.cli.Set(constant.RedisPrefixUser+user.Name, jd, constant.RedisExpirationUserLogin).Err()
}

func (s *serviceImpl) GetUserFromCache(name string) (*model.User, error) {
	val, err := s.cli.Get(constant.RedisPrefixUser + name).Bytes()
	if err != nil {
		return nil, err
	}
	user := model.User{}
	err = json.Unmarshal(val, &user)
	if err != nil {
		return nil, err
	}
	// 刷新 key 的过期时间
	s.cli.Expire(constant.RedisPrefixUser+name, constant.RedisExpirationUserLogin)
	return &user, nil
}

func (s *serviceImpl) DeleteUserFromCache(name string) error {
	return s.Del(constant.RedisPrefixUser + name)
}

func (s *serviceImpl) CheckTaskLimit(key string, limit int) bool {
	// 返回集合大小
	ret := s.cli.SCard(key).Val()
	if ret >= int64(limit) {
		return false
	}
	return true
}

func (s *serviceImpl) AddWorkingTaskIdToSet(key string, id string) error {
	if err := s.cli.SAdd(key, id).Err(); err != nil {
		return err
	}
	return nil
}

func (s *serviceImpl) DelWorkingTaskIdFromSet(key string, id string) error {
	if err := s.cli.SRem(key, id).Err(); err != nil {
		return err
	}
	return nil
}
