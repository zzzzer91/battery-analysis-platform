package server

import (
	"battery-analysis-platform/app/main/controller/api"
	"battery-analysis-platform/app/main/controller/auth"
	"battery-analysis-platform/app/main/controller/websocket"
	"battery-analysis-platform/app/main/middleware"
	"battery-analysis-platform/app/main/model"
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
		apiV1.GET("/sys-info", api.ShowSysInfo)

		apiV1.GET("/mining/base", api.ShowMiningBaseData)

		apiV1.POST("/mining/tasks", api.CreateMiningTask)
		apiV1.GET("/mining/tasks", api.ListMiningTask)
		apiV1.GET("/mining/tasks/:taskId", api.ShowMiningTaskData)
		apiV1.DELETE("/mining/tasks/:taskId", api.DeleteMiningTask)

		apiV1.POST("/dl/tasks", api.CreateDlTask)
		apiV1.GET("/dl/tasks", api.ListDlTask)
		apiV1.GET("/dl/tasks/:taskId", api.ShowDlTaskData)
		apiV1.DELETE("/dl/tasks/:taskId", api.DeleteDlTask)
	}
	apiV1NeedAuth := r.Group("/api/v1")
	apiV1.Use(middleware.PermissionRequired(model.UserTypeSuperUser))
	{
		apiV1NeedAuth.GET("/users", api.ListUser)
		apiV1NeedAuth.POST("/users", api.CreateUser)
		apiV1NeedAuth.PUT("/users/:name", api.ModifyUser)
	}

	wsV1 := r.Group("/websocket/v1")
	wsV1.Use(middleware.PermissionRequired(model.UserTypeCommonUser))
	{
		//wsV1.GET("/sys-info", websocket.ShowSysInfo)
		wsV1.GET("/mining/tasks", websocket.ListMiningTask)
	}
}
