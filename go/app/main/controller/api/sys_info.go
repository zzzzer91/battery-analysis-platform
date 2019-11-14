package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ShowSysInfo(c *gin.Context) {
	var s service.SysInfoShowService
	controller.GinResponse(c, &s)
}
