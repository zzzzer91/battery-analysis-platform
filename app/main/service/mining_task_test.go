package service

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTaskList(t *testing.T) {
	ast := assert.New(t)
	ret, err := GetTaskList()
	ast.Nil(err)
	s, _ := json.Marshal(ret)
	t.Log(string(s))
}

func TestDeleteTask(t *testing.T) {
	ast := assert.New(t)
	count, err := DeleteTask("none")
	ast.Nil(err)
	ast.Equal(int64(0), count)
}
