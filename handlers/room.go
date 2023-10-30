package handlers

import (
	"log"

	"github.com/gofiber/websocket/v2"
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
	}()

	clients[conn] = true
	for {
		var messageIn MessageIn
		err := conn.ReadJSON(&messageIn)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			delete(clients, conn)
			break
		}
		if messageIn.Message == `` {
			break
		}

		messageOut := MessageOut{
			Username: "Habi",
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
				log.Printf("Error occured: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
