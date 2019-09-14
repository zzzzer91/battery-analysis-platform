package api

import (
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var s service.MiningCreateTaskService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	data, err := s.CreateTask()
	code, msg := jd.HandleError(err)
	if code == jd.SUCCESS {
		msg = "创建成功"
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func GetTaskList(c *gin.Context) {
	data, err := service.GetTaskList()
	code, msg := jd.HandleError(err)
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func GetTask(c *gin.Context) {
	data, err := service.GetTask(c.Param("taskId"))
	code, msg := jd.HandleError(err)
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	_, err := service.DeleteTask(c.Param("taskId"))
	code, msg := jd.HandleError(err)
	if code == jd.SUCCESS {
		msg = "删除成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}
