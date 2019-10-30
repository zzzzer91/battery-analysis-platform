package websocket

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListDlTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		db.Redis.BRPop(taskWaitSigTimeout, "deeplearningTask:sigList")

		var s service.DlTaskListService
		res, err := s.Do()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = conn.WriteJSON(res); err != nil {
			fmt.Println(err)
			return
		}
	}
}
