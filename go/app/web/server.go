package web

import (
	"battery-analysis-platform/app/web/cache"
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/controller"
	"battery-analysis-platform/app/web/dal"
	"battery-analysis-platform/app/web/middleware"
	"battery-analysis-platform/app/web/producer"
	"github.com/gin-gonic/gin"
)

func Run() error {
	gin.SetMode(conf.App.Gin.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(conf.App.Gin.SecretKey))
	controller.Register(r)
	dal.Init()
	cache.Init()
	producer.Init()
	return r.Run(conf.App.Gin.HttpAddr)
}
