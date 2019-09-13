package service

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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

//func TestGetTask(t *testing.T) {
//	ret, err := GetTask("232eb995-e866-43ac-801b-4ea1bf35b5b6")
//	if err != nil {
//		t.Error(err)
//	}
//	s, _ := json.Marshal(ret)
//	t.Log(string(s))
//}

func TestDeleteTask(t *testing.T) {
	ast := assert.New(t)
	count, err := DeleteTask("none")
	if err != nil {
		t.Error(err)
	}
	ast.Equal(int64(0), count)
}
