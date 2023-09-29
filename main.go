package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"chat-app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	wsConf := websocket.Config{
		HandshakeTimeout: 100 * time.Second,
		Origins: []string{
			fmt.Sprintf("http://localhost%s", *addr),
			fmt.Sprintf("http://127.0.0.1%s", *addr),
		},
	}

	room := models.NewRoom()
	app.Use("/room", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Status(fiber.StatusUpgradeRequired).Send(nil)
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/room", websocket.New(room.Handler, wsConf))

	go room.Run()
	log.Fatal(app.Listen(*addr))
}
