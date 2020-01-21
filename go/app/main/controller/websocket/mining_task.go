package websocket

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ListMiningTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer conn.Close()

	s := service.MiningTaskListService{}

	closed := monitorWsClosed(conn)
	for {
		select {
		case <-closed:
			// 注意这里不能用 break，break只能跳出 select
			// 要用 return
			return
		default:
			res, err := s.Do()
			if err != nil {
				c.Error(err)
				return
			}
			if err = conn.WriteJSON(res); err != nil {
				c.Error(err)
				return
			}
		}
		// 等待数据改变
		// 注意：
		// 超时必须设置，不然前端关闭连接后，websocket会永久阻塞
		db.Redis.BLPop(consts.TaskWaitSigTimeout, "miningTask:sigList")
	}
}
