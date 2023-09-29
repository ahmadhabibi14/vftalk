package domain

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

// type Room struct {
// 	Clients map[*Client]bool
// 	Join    chan *Client
// 	Leave   chan *Client
// 	Forward chan []byte
// }

// func NewRoom() *Room {
// 	return &Room{
// 		Forward: make(chan []byte),
// 		Join:    make(chan *Client),
// 		Leave:   make(chan *Client),
// 		Clients: make(map[*Client]bool),
// 	}
// }

func RoomUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		c.Locals("ClientID", uuid.String())
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func RoomHandler(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}

// func (r *Room) Run() {
// 	for {
// 		select {
// 		case client := <-r.Join:
// 			r.Clients[client] = true
// 		case client := <-r.Leave:
// 			delete(r.Clients, client)
// 			close(client.Receive)
// 		case msg := <-r.Forward:
// 			for client := range r.Clients {
// 				client.Receive <- msg
// 			}
// 		}
// 	}
// }

// const (
// 	socketBufferSize  = 1824
// 	messageBufferSize = 256
// )

// func (r *Room) RoomHandler(conn *websocket.Conn) {
// 	client := &Client{
// 		Socket:  conn,
// 		Receive: make(chan []byte, messageBufferSize),
// 		Room:    r,
// 	}
// 	r.Join <- client

// 	go client.Write()
// 	client.Read()

// 	defer func() { r.Leave <- client }()
// }
