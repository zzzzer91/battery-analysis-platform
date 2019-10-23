package websocket

import (
	"battery-analysis-platform/app/main/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func ListMiningTask(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// read 指定时间后超时返回
		err := conn.SetReadDeadline(time.Now().Add(6 * time.Second))
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, _, err = conn.ReadMessage(); err != nil {
			fmt.Println(err)
			return
		}
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
		// 服务端也做速度控制
		time.Sleep(time.Second * 3)
	}
}
