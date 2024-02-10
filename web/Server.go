package web

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"vftalk/configs"
	"vftalk/handlers/apis"
	"vftalk/handlers/page"
	"vftalk/middlewares"
	"vftalk/models/mailer"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
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
	mlr := mailer.NewMailer(w.Log)
	db, err := configs.ConnectMariaDB()
	if err != nil {
		w.Log.Error().Str("error", err.Error()).Msg("failed when try to connect database")
	}
	oauth := configs.EnvOAuth()

	engine := html.New("./views/pages/dist", ".html")
	app := fiber.New(fiber.Config{
		AppName: w.AppName,
		Views:   engine,
		Prefork: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var code int = fiber.StatusNotFound
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			c.Status(fiber.StatusNotFound)
			return c.Render("404", fiber.Map{
				"Title":   fmt.Sprintf("%d - %s", code, http.StatusText(code)),
				"Code":    code,
				"Status":  http.StatusText(code),
				"Message": http.StatusText(code),
			})
		},
	})

	app.Use(favicon.New(favicon.Config{
		File: "./views/public/favicons/favicon.ico",
		URL:  "/favicon.ico",
	}))
	app.Use(requestid.New())
	app.Use(logger.New(middlewares.NewFiberLogger()))
	app.Use(limiter.New(middlewares.Limiter))
	app.Use(cors.New(middlewares.CORSConfig))
	app.Use(recover.New())
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./docs/apidocs.json",
		Path:     "docs",
		Title:    "VFtalk | API Docs",
		CacheAge: int(30 * time.Minute),
	}))

	app.Static("/", "./views/public")
	app.Static("/", "./contents")
	app.Static("_astro", "./views/pages/dist/_astro")

	apiHandler := &apis.ApisHandler{
		Mailer: mlr,
		Log:    w.Log,
		Db:     db,
		OAuth:  oauth,
	}
	pageHandler := &page.PageHandler{
		Log:    w.Log,
		Db:     db,
		OAuth:  oauth,
		Domain: os.Getenv("WEB_DOMAIN"),
	}

	WebViews(app, pageHandler)
	ApiRoutes(app, apiHandler)
	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}
