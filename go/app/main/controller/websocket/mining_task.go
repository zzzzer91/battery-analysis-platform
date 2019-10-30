package websocket

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	taskWaitSigTimeout = time.Second * 3
)

func ListMiningTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// TODO 需要设法读取对端关闭

		// 等待数据改变
		// 注意：
		// 超时必须设置，不然前端关闭连接后，websocket会永久阻塞
		db.Redis.BRPop(taskWaitSigTimeout, "miningTask:sigList")

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
