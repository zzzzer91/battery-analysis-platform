package server

import (
	"battery-anlysis-platform/app/main/middleware"
	"battery-anlysis-platform/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Start() error {
	gin.SetMode(conf.App.Main.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(conf.App.Main.SecretKey))
	register(r)

	return r.Run(conf.App.Main.HttpAddr)
}
