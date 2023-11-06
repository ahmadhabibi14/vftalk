package presentation

import (
	"time"

	"vftalk/handlers"
	"vftalk/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	wsConf := websocket.Config{
		HandshakeTimeout: 800 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}

	api := app.Group("/api")

	api.Use("/room", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	api.Get("/room", websocket.New(handlers.HandleClients, wsConf))

	api.Post("/login", middlewares.IsLoggedIn, handlers.Login)
	api.Post("/register", middlewares.IsLoggedIn, handlers.Register)
}
