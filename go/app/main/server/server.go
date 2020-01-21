package server

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/app/main/middleware"
	"github.com/gin-gonic/gin"
)

func Start() error {
	gin.SetMode(conf.App.Gin.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(conf.App.Gin.SecretKey))
	register(r)
	return r.Run(conf.App.Gin.HttpAddr)
}
