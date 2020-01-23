package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func GetSysInfo(c *gin.Context) {
	s := service.GetSysInfoService{}
	controller.JsonResponse(c, &s)
}
