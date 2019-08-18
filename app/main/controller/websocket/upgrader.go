package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func upgradeHttpConn(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	conn, err := wsUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}
