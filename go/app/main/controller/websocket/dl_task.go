package websocket

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ListDlTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer conn.Close()

	s := service.DlTaskListService{}

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
		db.Redis.BLPop(consts.WsSendInterval, consts.RedisKeyDlTaskSigList)
	}
}

func ShowDlTaskTraningHistory(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer conn.Close()

	s := service.DlTaskShowTraningHistoryService{
		Id:            c.Param("taskId"),
		ReadFromRedis: true,
	}

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

		prefixStr := consts.RedisPrefixDlTaskTrainingHistory + s.Id + ":"
		db.Redis.BLPop(
			consts.WsSendInterval,
			prefixStr+consts.RedisCommonKeySigList)
	}
}
