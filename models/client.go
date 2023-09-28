package models

import "github.com/gofiber/websocket/v2"

type Client struct {
	Socket  *websocket.Conn
	Receive chan []byte
	Room    *Room
}

func (c *Client) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		c.Room.Forward <- msg
	}
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
