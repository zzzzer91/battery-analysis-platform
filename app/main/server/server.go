package server

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/app/main/middleware"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	ginConf := &conf.App.Gin
	gin.SetMode(ginConf.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(ginConf.SecretKey))
	register(r)
	return r
}

func Start() error {
	r := Default()
	return r.Run(conf.App.Gin.HttpAddr)
}
