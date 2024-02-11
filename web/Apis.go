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

	// Auth
	api.Post("/register", middlewares.ContentJSON, apis.AuthRegister)
	api.Post("/login", middlewares.ContentJSON, apis.AuthLogin)
	api.Get("/oauth/google", apis.OAuthGoogle)
	api.Post("logout", apis.AuthLogout)

	// Users Specific Data
	api.Put("/user-update-profile", middlewares.AuthJWT, middlewares.ContentJSON, apis.UpdateProfile)
	api.Put("/user-update-avatar", middlewares.AuthJWT, middlewares.ContentMultipartForm, apis.UpdateAvatar)
	api.Post("/user-lists", middlewares.AuthJWT, middlewares.ContentJSON, apis.UserLists)

	// Chat Rooms
	api.Get("/room-general", middlewares.AuthJWT, middlewares.Websocket, websocket.New(apis.ChatRoomGeneral, WEBSOCKET_CONF))

	// Debug
	api.Get("/debug", apis.Debug)
}
