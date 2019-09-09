package conf

import "battery-anlysis-platform/pkg/conf"

type confParams struct {
	RunMode   string
	HttpAddr  string
	MysqlUri  string
	MongoUri  string
	RedisUri  string
	SecretKey string
}

var Params confParams

func init() {
	err := conf.Load("conf/app.ini", "main", &Params)
	if err != nil {
		panic(err)
	}
}
