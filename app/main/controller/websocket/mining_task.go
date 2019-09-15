package websocket

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ListMiningTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	closed := wsClientClosed(conn)
	for {
		select {
		case <-closed:
			return
		default:
			var s service.MiningTaskListService
			res, err := s.Do()
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteJSON(res); err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second * 3)
		}
	}
}
