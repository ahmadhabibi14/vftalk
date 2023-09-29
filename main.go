package main

import (
	"log"
	"time"

	"chat-app/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	wsConf := websocket.Config{
		HandshakeTimeout: 100 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}

	// room := domain.NewRoom()
	app.Use("/room", domain.RoomUpgrade)
	app.Get("/room", websocket.New(domain.RoomHandler, wsConf))

	// go room.Run()
	log.Fatal(app.Listen(":8080"))
}
