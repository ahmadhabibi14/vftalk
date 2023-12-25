package app

import (
	"time"

	"vftalk/handlers"
	"vftalk/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App, h handlers.Handler) {
	wsConf := websocket.Config{
		HandshakeTimeout: 800 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}
	api := app.Group("/api")

	api.Get("/room", middlewares.AuthJWT, middlewares.Websocket, websocket.New(handlers.HandleClients, wsConf))
	api.Post("/login", handlers.Login)
	api.Post("/register", h.Register)
	api.Get("/oauth/google", handlers.OAuthGoogle)
	api.Post("/userdata", handlers.GetUserData)
	api.Post("/user-update-avatar", middlewares.AuthJWT, handlers.UpdateProfilePicture)
	api.Post("/user-update-profile", middlewares.AuthJWT, handlers.UpdateProfile)

	api.Get("/debug", h.Debug)
}
