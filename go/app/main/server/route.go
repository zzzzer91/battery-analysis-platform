package server

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/controller/api"
	"battery-analysis-platform/app/main/controller/auth"
	"battery-analysis-platform/app/main/controller/file"
	"battery-analysis-platform/app/main/controller/websocket"
	"battery-analysis-platform/app/main/middleware"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	r.GET("/login", auth.Login)
	r.POST("/login", auth.Login)

	rootPath := r.Group("/")
	rootPath.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		rootPath.POST("/logout", auth.Logout)
	}

	apiV1Path := r.Group("/api/v1")
	apiV1Path.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		apiV1Path.POST("/self/change-password", api.ChangePassword)

		apiV1Path.GET("/sys-info", api.ShowSysInfo)

		apiV1Path.GET("/mining/base", api.ShowMiningBaseData)

		apiV1Path.POST("/mining/tasks", api.CreateMiningTask)
		apiV1Path.GET("/mining/tasks", api.ListMiningTask)
		apiV1Path.GET("/mining/tasks/:taskId/data", api.ShowMiningTaskData)
		apiV1Path.DELETE("/mining/tasks/:taskId", api.DeleteMiningTask)

		apiV1Path.POST("/dl/tasks", api.CreateDlTask)
		apiV1Path.GET("/dl/tasks", api.ListDlTask)
		apiV1Path.GET("/dl/tasks/:taskId/training-history", api.ShowDlTaskTraningHistory)
		apiV1Path.GET("/dl/tasks/:taskId/eval-result", api.ShowDlEvalResultHistory)
		apiV1Path.DELETE("/dl/tasks/:taskId", api.DeleteDlTask)
	}

	apiV1NeedPermissionPath := r.Group("/api/v1")
	apiV1NeedPermissionPath.Use(middleware.PermissionRequired(consts.UserTypeSuperUser))
	{
		apiV1NeedPermissionPath.GET("/users", api.ListUser)
		apiV1NeedPermissionPath.POST("/users", api.CreateUser)
		apiV1NeedPermissionPath.PUT("/users/:name", api.ModifyUser)
	}

	wsV1Path := r.Group("/websocket/v1")
	wsV1Path.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		//wsV1Path.GET("/sys-info", websocket.ShowSysInfo)
		wsV1Path.GET("/mining/tasks", websocket.ListMiningTask)
		wsV1Path.GET("/dl/tasks", websocket.ListDlTask)
		wsV1Path.GET("/dl/tasks/:taskId/training-history", websocket.ShowDlTaskTraningHistory)
	}

	filePath := r.Group("/file")
	filePath.Use(middleware.PermissionRequired(consts.UserTypeCommonUser))
	{
		filePath.GET("/dl/model/:taskId", file.DownloadDlModel)
	}
}
