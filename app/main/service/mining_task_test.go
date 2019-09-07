package service

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/app/main/dao"
	gconf "battery-anlysis-platform/pkg/conf"
	"encoding/json"
	"testing"
)

func init() {
	confParams := &conf.Params{}
	gconf.Load("conf/app.ini", "main", confParams)
	dao.InitMongo(confParams.MongoUri)
}

func TestGetTaskList(t *testing.T) {
	ret, err := GetTaskList()
	if err != nil {
		t.Error(err)
	}
	s, _ := json.Marshal(ret)
	t.Log(string(s))
}

func TestGetTask(t *testing.T) {
	ret, err := GetTask("232eb995-e866-43ac-801b-4ea1bf35b5b6")
	if err != nil {
		t.Error(err)
	}
	s, _ := json.Marshal(ret)
	t.Log(string(s))
}
