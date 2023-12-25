package app

import (
	"log"

	"vftalk/conf"
	"vftalk/handlers"
	"vftalk/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/handlebars/v2"
)

type WebServer struct {
	AppName string
	Cfg     conf.WebConf
	handlers.Handler
}

func (w *WebServer) Start() {
	engine := handlebars.New("./views/routes", ".hbs")
	app := fiber.New(fiber.Config{
		AppName: w.AppName,
		Views:   engine,
		Prefork: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusNotFound)
			return c.Render("404", fiber.Map{
				"Title":   "404 - Page not found",
				"Message": "Page not found",
			})
		},
	})
	app.Use(favicon.New(favicon.Config{
		File: "./views/public/favicons/favicon.ico",
		URL:  "/favicon.ico",
	}))
	app.Use(requestid.New())
	app.Use(logger.New(middlewares.LoggerConfig))
	app.Use(limiter.New(middlewares.Limiter))
	app.Use(cors.New(middlewares.CORSConfig))

	app.Static("/static", "./views/static")
	app.Static("/public", "./views/public")
	app.Static("/files", "./uploads")

	WebViews(app)
	ApiRoutes(app, w.Handler)

	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}
