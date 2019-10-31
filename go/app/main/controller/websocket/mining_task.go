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

	closed := monitorWsClosed(conn)
	for {
		// 等待数据改变
		// 注意：
		// 超时必须设置，不然前端关闭连接后，websocket会永久阻塞
		db.Redis.BRPop(taskWaitSigTimeout, "miningTask:sigList")

		select {
		case <-closed:
			// 注意这里不能用 break，break只能跳出 select
			// 要用 return
			return
		default:
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
}
