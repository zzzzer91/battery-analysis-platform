package service

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/app/main/dao"
	gconf "battery-anlysis-platform/pkg/conf"
	"testing"
)

func init() {
	confParams := &conf.Params{}
	gconf.Load("conf/app.ini", "main", confParams)
	dao.InitMySQL(confParams.MysqlUri)
}

func TestGetUsers(t *testing.T) {
	ret, _ := GetUsers()
	t.Log(ret)
}
