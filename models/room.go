package models

import (
	"github.com/gofiber/websocket/v2"
)

type Room struct {
	Clients map[*Client]bool
	Join    chan *Client
	Leave   chan *Client
	Forward chan []byte
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan []byte),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

// func (r *Room) Upgrade(c *fiber.Ctx) error {
// 	if websocket.IsWebSocketUpgrade(c) {
// 		c.Locals("allowed", true)
// 		return c.Next()
// 	}
// 	return fiber.ErrUpgradeRequired
// }

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Receive)
		case msg := <-r.Forward:
			for client := range r.Clients {
				client.Receive <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1824
	messageBufferSize = 256
)

func (r *Room) Handler(conn *websocket.Conn) {
	client := &Client{
		Socket:  conn,
		Receive: make(chan []byte, messageBufferSize),
		Room:    r,
	}
	r.Join <- client
	defer func() { r.Leave <- client }()
	go client.Write()
	client.Read()
}
