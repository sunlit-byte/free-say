package room

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	chatRoom *ChatRoom
	ch       chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.chatRoom.clients[c] = false
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.chatRoom.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message, ok := <-c.ch
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			break
		}
		c.conn.WriteMessage(websocket.TextMessage, message)
	}
}
