package server

import (
	"battery-anlysis-platform/app/main/middleware"
	"battery-anlysis-platform/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Start(ginConf *conf.GinConf) error {
	gin.SetMode(ginConf.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(ginConf.SecretKey))
	register(r)

	return r.Run(ginConf.HttpAddr)
}
