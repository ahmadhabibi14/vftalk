package domain

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"chat-app/utils"
)

type Room struct {
	sync.RWMutex
	Clients map[*Client]bool
	Join    chan *Client
	Leave   chan *Client
	Forward chan []byte
}

func NewRoom() *Room {
	r := &Room{
		Forward: make(chan []byte),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
	go r.Run()
	return r
}

func RoomUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("ClientID", utils.GenerateRandomID(10))
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			fmt.Println("Join")
			r.Clients[client] = true
		case client := <-r.Leave:
			fmt.Println("Leave")
			delete(r.Clients, client)
			close(client.Receive)
		case msg := <-r.Forward:
			for client := range r.Clients {
				fmt.Println("Broadcast")
				client.Receive <- msg
			}
		}
	}
}

func (r *Room) RoomHandler(conn *websocket.Conn) {
	client := &Client{
		Socket:  conn,
		Receive: make(chan []byte, 256),
		Room:    r,
	}
	r.Join <- client

	go client.Write()
	client.Read()

	defer func() { r.Leave <- client }()
}
