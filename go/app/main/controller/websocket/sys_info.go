// 返回系统状态，内存，CPU 等

package websocket

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
	"time"
)

func ShowSysInfo(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	for {
		s := service.SysInfoShowService{}
		res, err := s.Do()
		if err != nil {
			c.Error(err)
			return
		}
		if err := conn.WriteJSON(res); err != nil {
			c.Error(err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}
