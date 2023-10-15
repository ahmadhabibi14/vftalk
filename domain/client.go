package domain

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

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
		fmt.Println(string(msg))
		c.Room.Forward <- msg
	}
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for {
		select {
		case msg, ok := <-c.Receive:
			if !ok {
				if err := c.Socket.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection closed: ", err)
				}
				return
			}
			if err := c.Socket.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
			}
		}
	}
}
