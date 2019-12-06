package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateDlTask(c *gin.Context) {
	s := service.DlTaskCreateService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func DeleteDlTask(c *gin.Context) {
	s := service.DlTaskDeleteService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}

func ListDlTask(c *gin.Context) {
	s := service.DlTaskListService{}
	controller.JsonResponse(c, &s)
}

func ShowDlTaskTraningHistory(c *gin.Context) {
	s := service.DlTaskShowTraningHistoryService{
		Id:            c.Param("taskId"),
		ReadFromRedis: false,
	}
	controller.JsonResponse(c, &s)
}

func ShowDlEvalResultHistory(c *gin.Context) {
	s := service.DlTaskShowEvalResultService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}
