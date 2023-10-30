package presentation

import (
	"time"

	"vftalk/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ApiRoutes(app *fiber.App) {
	wsConf := websocket.Config{
		HandshakeTimeout: 800 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}

	api := app.Group("/api")

	api.Use("/room", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	api.Get("/room", websocket.New(handlers.HandleClients, wsConf))
}
