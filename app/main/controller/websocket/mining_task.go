package websocket

import (
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetTaskList(c *gin.Context) {
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
			data, err := service.GetTaskList()
			code, msg := jd.HandleError(err)
			res := jd.Build(code, msg, data)
			if err := conn.WriteJSON(res); err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second * 3)
		}
	}
}
