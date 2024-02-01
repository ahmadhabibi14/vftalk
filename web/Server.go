package web

import (
	"log"

	"vftalk/configs"
	"vftalk/handlers/apis"
	"vftalk/handlers/page"
	"vftalk/middlewares"
	"vftalk/models/mailer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/rs/zerolog"
)

type WebServer struct {
	AppName string
	Cfg     configs.WebConf
	Log     *zerolog.Logger
}

func NewWebServer(cfg configs.WebConf, lg *zerolog.Logger) *WebServer {
	return &WebServer{
		AppName: "VFtalk - Chat App",
		Cfg:     cfg,
		Log:     lg,
	}
}

func (w *WebServer) Start() {
	mlr := MailServer(w.Log)
	db := configs.ConnectMariaDB()
	oauth := configs.EnvOAuth()

	engine := handlebars.New("./views/pages", ".hbs")
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
	app.Use(recover.New())

	app.Static("/static", "./views/static")
	app.Static("/public", "./views/public")
	app.Static("/files", "./uploads")

	apiHandler := &apis.ApisHandler{
		Mailer: mlr,
		Log:    w.Log,
		Db:     db,
		OAuth:  oauth,
	}
	pageHandler := &page.PageHandler{
		Log:   w.Log,
		Db:    db,
		OAuth: oauth,
	}

	WebViews(app, pageHandler)
	ApiRoutes(app, apiHandler)
	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}

func MailServer(l *zerolog.Logger) mailer.Mailer {
	mlr := mailer.Mailer{}
	mh, err := mailer.NewMailhog(configs.EnvMailhog())
	if err != nil {
		l.Error().Str(`Error: `, err.Error()).Msg(`Cannot load mailhog`)
	}
	mlr.SendMailFunc = mh.SendEmail
	return mlr
}
