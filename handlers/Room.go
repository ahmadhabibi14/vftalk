package handlers

import (
	"fmt"
	"vftalk/conf"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

type (
	MessageIn struct {
		Message string `json:"message"`
	}
	MessageOut struct {
		Username string `json:"username"`
		Message  string `json:"message"`
	}
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan MessageOut)
)

func HandleClients(conn *websocket.Conn) {
	go broadcastMessagesToClients()
	defer func() {
		conn.Close()
		log.Error(`User disconnected`)
	}()

	clients[conn] = true

	u, _ := conf.WsGetUsernameFromJWT(conn)
	username := fmt.Sprintf("%v", u)

	for {
		var messageIn MessageIn
		err := conn.ReadJSON(&messageIn)
		if err != nil {
			delete(clients, conn)
			break
		}
		if messageIn.Message == `` {
			break
		}

		messageOut := MessageOut{
			Username: username,
			Message:  messageIn.Message,
		}
		broadcast <- messageOut
	}
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
