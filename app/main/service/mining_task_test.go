package service

import (
	"encoding/json"
	"testing"
)

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
