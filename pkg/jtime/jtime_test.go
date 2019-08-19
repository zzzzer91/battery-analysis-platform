package jtime

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJSONTime(t *testing.T) {
	ast := assert.New(t)
	t1 := time.Date(2019, 2, 2, 3, 46, 29, 0, time.Local)
	testCase := &JSONTime{Time: t1}
	expected := `"2019-02-02 03:46:29"`
	ret, err := json.Marshal(testCase)
	if err != nil {
		t.Log(err)
		//  标记测试函数为失败，然后继续执行（剩下的测试）。
		t.Fail()
	}
	ast.Equal(expected, string(ret))
}
