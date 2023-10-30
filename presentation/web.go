package presentation

import (
	"log"

	"vftalk/conf"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/handlebars/v2"
)

type WebServer struct {
	AppName string
	Cfg     conf.WebConf
}

func (w *WebServer) Start() {
	engine := handlebars.New("./views/routes", ".hbs")
	app := fiber.New(fiber.Config{
		AppName: w.AppName,
		Views:   engine,
		Prefork: false,
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 03:04 PM",
		TimeZone:   "Asia/Makassar",
	}))

	app.Static("/static", "./views/static")
	app.Static("/public", "./views/public")

	WebViews(app)
	ApiRoutes(app)

	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}
