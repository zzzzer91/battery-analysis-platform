package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetSysInfo(c *gin.Context) {
	res := jd.Build(jd.SUCCESS, "", service.GetSysInfo())
	c.JSON(200, res)
}
