package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetBasicData(c *gin.Context) {
	var code int
	var msg string
	var data interface{}

	s := &service.MiningBaseService{
		DataComeFrom: c.Query("dataComeFrom"),
		StartDate:    c.Query("startDate"),
		NeedParams:   c.Query("needParams"),
		DataLimit:    c.Query("dataLimit"),
	}

	ret, err := s.Query()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "查询成功"
		data = ret
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
