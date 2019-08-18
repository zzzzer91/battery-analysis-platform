// 返回系统状态，内存，CPU 等

package websocket

import (
	"battery-anlysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetSysInfo(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		data := service.GetSysInfo()
		if err := conn.WriteJSON(data); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}
