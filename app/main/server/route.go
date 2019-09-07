package server

import (
	"battery-anlysis-platform/app/main/controller/api"
	"battery-anlysis-platform/app/main/controller/auth"
	"battery-anlysis-platform/app/main/controller/websocket"
	"battery-anlysis-platform/app/main/middleware"
	"battery-anlysis-platform/app/main/model"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	r.GET("/login", auth.Login)
	r.POST("/login", auth.Login)

	root := r.Group("/")
	root.Use(middleware.PermissionRequired(model.UserTypeCommonUser))
	{
		root.POST("/logout", auth.Logout)
	}

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.PermissionRequired(model.UserTypeCommonUser))
	{
		apiV1.GET("/ping", api.Pong)

		apiV1.GET("/sys-info", api.GetSysInfo)

		apiV1.GET("/mining/base", api.GetBasicData)

		apiV1.GET("/mining/tasks", api.GetTaskList)

		apiV1.GET("/mining/tasks/:taskId", api.GetTask)
	}
	apiV1NeedAuth := r.Group("/api/v1")
	apiV1.Use(middleware.PermissionRequired(model.UserTypeSuperUser))
	{
		apiV1NeedAuth.GET("/users", api.GetUsers)
		apiV1NeedAuth.POST("/users", api.CreateUser)
		apiV1NeedAuth.PUT("/users/:name", api.ModifyUser)
	}

	wsV1 := r.Group("/websocket/v1")
	wsV1.Use(middleware.PermissionRequired(model.UserTypeCommonUser))
	{
		wsV1.GET("/echo", websocket.Echo)
		wsV1.GET("/sys-info", websocket.GetSysInfo)
	}
}
