package redis

import (
	"battery-analysis-platform/app/web/model"
	"github.com/go-redis/redis/v7"
	"time"
)

type Service interface {
	Del(strs ...string) error
	LRange(key string, start, stop int64) ([]string, error)
	BLPop(timeout time.Duration, keys ...string) ([]string, error)
	AddUserToCache(user *model.User) error
	GetUserFromCache(name string) (*model.User, error)
	DeleteUserFromCache(name string) error
	CheckTaskLimit(key string, limit int) bool
	AddWorkingTaskIdToSet(key string, id string) error
	DelWorkingTaskIdFromSet(key string, id string) error
}

func NewService(cli *redis.Client) Service {
	return &serviceImpl{cli}
}
