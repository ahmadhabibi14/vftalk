package web

import (
	"time"
	"vftalk/handlers/apis"
	"vftalk/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var WEBSOCKET_CONF = websocket.Config{
	HandshakeTimeout: 800 * time.Second,
	ReadBufferSize:   1824,
	WriteBufferSize:  256,
}

func ApiRoutes(app *fiber.App, apis *apis.ApisHandler) {
	api := app.Group("/api")
	api.Post("/register", apis.AuthRegister)
	api.Get("/oauth/google", apis.OAuthGoogle)
	api.Get("/room", middlewares.AuthJWT, middlewares.Websocket, websocket.New(apis.UserChatRoom, WEBSOCKET_CONF))
	// api.Post("/login", handlers.Login)

	// api.Post("/userdata", handlers.GetUserData)
	// api.Post("/user-update-avatar", middlewares.AuthJWT, handlers.UpdateProfilePicture)
	// api.Post("/user-update-profile", middlewares.AuthJWT, handlers.UpdateProfile)

	api.Get("/debug", apis.Debug)
}
