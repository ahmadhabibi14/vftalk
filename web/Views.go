package web

import (
	"vftalk/handlers/page"
	"vftalk/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WebViews(app *fiber.App, page *page.PageHandler) {
	app.Get("/", page.Index)
	app.Get("/register", middlewares.IsLoggedIn, page.Register)
	app.Get("/login", middlewares.IsLoggedIn, page.Login)
	app.Get("/oauth/google", middlewares.IsLoggedIn, page.OAuthGoogle)
	app.Get("/about", page.About)
	app.Get("/contact", page.Contact)

	// app.Get("/explore", middlewares.AuthJWT, page.Explore)
	// app.Get("/profile", middlewares.AuthJWT, page.Profile)
	// app.Get("/direct", middlewares.AuthJWT, page.DirectChat)
	app.Get("/.env", page.RickRoll)
}
