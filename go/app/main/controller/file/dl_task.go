package file

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func DownloadDlModel(c *gin.Context) {
	var s service.DlDownloadModelService
	s.Id = c.Param("taskId")
	controller.FileResponse(c, &s)
}
