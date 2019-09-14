package api

import (
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetSysInfo(c *gin.Context) {
	data, err := service.GetSysInfo()
	code, msg := jd.HandleError(err)
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
