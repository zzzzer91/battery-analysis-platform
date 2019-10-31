package api

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateDlTask(c *gin.Context) {
	var s service.DlTaskCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func DeleteDlTask(c *gin.Context) {
	var s service.DlTaskDeleteService
	s.Id = c.Param("taskId")
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ListDlTask(c *gin.Context) {
	var s service.DlTaskListService
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ShowDlTaskTraningHistory(c *gin.Context) {
	var s service.DlTaskShowTraningHistoryService
	s.Id = c.Param("taskId")
	s.ReadFromRedis = false
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ShowDlEvalResultHistory(c *gin.Context) {
	var s service.DlTaskShowEvalResultService
	s.Id = c.Param("taskId")
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}
