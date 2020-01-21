package consts

import "time"

// redis key 前缀
const (
	redisProjectKeyPrefix = "ubattery:"
	RedisUserKeyPrefix    = redisProjectKeyPrefix + "users:"
)

const (
	// redis 中用户登录状态过期时间
	RedisUserLoginExpiration = time.Hour * 3
)
