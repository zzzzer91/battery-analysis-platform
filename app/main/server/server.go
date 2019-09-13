package server

import (
	"battery-anlysis-platform/app/main/middleware"
	"battery-anlysis-platform/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	ginConf := &conf.App.Main.Gin
	gin.SetMode(ginConf.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(ginConf.SecretKey))
	register(r)
	return r
}

func Start() error {
	r := Default()
	return r.Run(conf.App.Main.Gin.HttpAddr)
}
