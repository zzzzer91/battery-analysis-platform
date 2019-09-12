package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrader *websocket.Upgrader

func init() {
	wsUpgrader = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		// 取消 ws 跨域校验，http 头部的 origin 不对，会返回 403，这里禁止这个功能
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}

func upgradeHttpConn(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	conn, err := wsUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}
