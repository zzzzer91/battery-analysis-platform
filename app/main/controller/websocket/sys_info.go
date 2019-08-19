// 返回系统状态，内存，CPU 等

package websocket

import (
	"battery-anlysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
	"time"
)

func GetSysInfo(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	for {
		data, _ := service.GetSysInfo()
		if err := conn.WriteJSON(data); err != nil {
			c.AbortWithError(500, err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}
