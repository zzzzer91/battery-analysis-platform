package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetTaskList(c *gin.Context) {
	var code int
	var msg string
	var data interface{}
	taskList, err := service.GetTaskList()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		data = taskList
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func GetTask(c *gin.Context) {
	var code int
	var msg string
	var data interface{}
	taskList, err := service.GetTask(c.Param("taskId"))
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		data = taskList
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
