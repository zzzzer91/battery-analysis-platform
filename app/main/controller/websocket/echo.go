// 用于测试 websocket

package websocket

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Echo(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err = conn.WriteMessage(t, msg); err != nil {
			log.Println(err)
			return
		}
	}
}
