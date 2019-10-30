package websocket

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListMiningTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 等待数据改变
		db.Redis.BRPop(model.TaskWaitSigTimeout, "miningTask:sigList")

		var s service.MiningTaskListService
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
