package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetSysInfo(c *gin.Context) {
	var code int
	var msg string
	data, err := service.GetSysInfo()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
