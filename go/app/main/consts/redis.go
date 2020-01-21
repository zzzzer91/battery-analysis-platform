package consts

import "time"

const (
	// redis 的键前缀
	RedisKeyPrefix = "ubattery:users:"
	// 缓存中 key 过期时间
	RedisKeyExpiration = time.Hour * 3
)
