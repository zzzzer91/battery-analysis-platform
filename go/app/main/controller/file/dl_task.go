package file

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func DownloadDlModel(c *gin.Context) {
	s := service.DownloadDlModelService{
		Id: c.Param("taskId"),
	}
	controller.FileResponse(c, &s)
}
