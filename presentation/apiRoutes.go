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

	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.Register)
	api.Get("/oauth/google", handlers.OAuthGoogle)
	api.Post("/userdata", handlers.GetUserData)
	api.Post("/user-active-list", handlers.GetUserActiveLists)
	api.Post("/user-update-active", handlers.UpdateUserLastActive)
	api.Post("/user-update-avatar", middlewares.AuthJWT, handlers.UpdateProfilePicture)
}
