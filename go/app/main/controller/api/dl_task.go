package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateDlTask(c *gin.Context) {
	var s service.DlTaskCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.GinResponse(c, &s)
}

func DeleteDlTask(c *gin.Context) {
	var s service.DlTaskDeleteService
	s.Id = c.Param("taskId")
	controller.GinResponse(c, &s)
}

func ListDlTask(c *gin.Context) {
	var s service.DlTaskListService
	controller.GinResponse(c, &s)
}

func ShowDlTaskTraningHistory(c *gin.Context) {
	var s service.DlTaskShowTraningHistoryService
	s.Id = c.Param("taskId")
	s.ReadFromRedis = false
	controller.GinResponse(c, &s)
}

func ShowDlEvalResultHistory(c *gin.Context) {
	var s service.DlTaskShowEvalResultService
	s.Id = c.Param("taskId")
	controller.GinResponse(c, &s)
}
