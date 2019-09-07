package server

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/app/main/middleware"
	"github.com/gin-gonic/gin"
)

func Start() error {
	gin.SetMode(conf.Params.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(conf.Params.SecretKey))
	register(r)

	return r.Run(conf.Params.HttpAddr)
}
