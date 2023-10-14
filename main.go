package main

import (
	"log"
	"runtime"
	"time"

	"chat-app/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/gofiber/websocket/v2"
)

func init() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)
}

func main() {
	engine := django.New("./views", ".django")
	app := fiber.New(fiber.Config{
		AppName: "Habi Chat App",
		Views:   engine,
		Prefork: true,
	})
	wsConf := websocket.Config{
		HandshakeTimeout: 800 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}
	app.Static("/static", "./views/static")

	room := domain.NewRoom()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Habi Chat App",
		})
	})
	app.Use("/room", domain.RoomUpgrade)
	app.Get("/room", websocket.New(room.RoomHandler, wsConf))
	go room.Run()
	log.Fatal(app.Listen(":8080"))
}
