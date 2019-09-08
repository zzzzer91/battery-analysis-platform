package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var s service.MiningCreateTaskService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	var code int
	var msg string
	data, err := s.CreateTask()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "创建成功"
	}

	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func GetTaskList(c *gin.Context) {
	var code int
	var msg string
	data, err := service.GetTaskList()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func GetTask(c *gin.Context) {
	var code int
	var msg string
	data, err := service.GetTask(c.Param("taskId"))
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	var code int
	var msg string
	_, err := service.DeleteTask(c.Param("taskId"))
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "删除成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}
