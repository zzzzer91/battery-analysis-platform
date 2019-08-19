// 用于测试 websocket

package websocket

import (
	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if err = conn.WriteMessage(t, msg); err != nil {
			c.AbortWithError(500, err)
			return
		}
	}
}
