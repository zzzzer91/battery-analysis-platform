package conf

import "battery-anlysis-platform/pkg/conf"

type confParams struct {
	RunMode   string
	HttpAddr  string
	MysqlUri  string
	MongoUri  string
	SecretKey string
}

var Params confParams

func init() {
	conf.Load("conf/app.ini", "main", &Params)
}
