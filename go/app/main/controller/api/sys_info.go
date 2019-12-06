package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ShowSysInfo(c *gin.Context) {
	s := service.SysInfoShowService{}
	controller.JsonResponse(c, &s)
}
