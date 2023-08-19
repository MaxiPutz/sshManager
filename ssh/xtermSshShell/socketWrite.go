package xtermsshshell

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type socketWrite struct {
	channel string
	socket  *websocket.Conn
}

type writeStruct struct {
	Msg string `json:"msg"`
	Id  string `json:"id"`
}

func (sw socketWrite) Write(p []byte) (n int, err error) {

	ws := writeStruct{
		Msg: string(p),
		Id:  sw.channel,
	}

	fmt.Printf("sw: %v\n", ws)

	b, _ := json.Marshal(&ws)

	sw.socket.WriteMessage(websocket.TextMessage, b)
	return len(p), nil
}
