package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEcho(t *testing.T) {
	ast := assert.New(t)

	url := "ws://localhost:8080/websocket/v1/echo"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatal(err)
	}
	writeMsg := []byte("hello, world!")
	if err := ws.WriteMessage(websocket.TextMessage, writeMsg); err != nil {
		t.Fatal(err)
	}
	msgType, readMsg, err := ws.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	ast.Equal(websocket.TextMessage, msgType)
	ast.Equal(writeMsg, readMsg)
}
