package room

import (
	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn) {
	client := &Client{
		conn:     conn,
		chatRoom: DefaultRoom,
		ch:       make(chan []byte, 256),
	}
	client.chatRoom.clients[client] = true
	go client.writePump()
	go client.readPump()

}
