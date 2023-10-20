package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cpu := runtime.NumCPU()
	log.Println(cpu)
	runtime.GOMAXPROCS(cpu)
}

type Message struct {
	Message string `json:"message"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

func HandleClients(conn *websocket.Conn) {
	go broadcastMessagesToClients()
	defer func() {
		conn.Close()
	}()

	clients[conn] = true
	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			delete(clients, conn)
			break
		}
		broadcast <- message
	}
}

func main() {
	engine := handlebars.New("./views/routes", ".hbs")
	app := fiber.New(fiber.Config{
		AppName: "VFtalk - Chat App",
		Views:   engine,
		Prefork: false,
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 03:04 PM",
		TimeZone:   "Asia/Makassar",
	}))
	wsConf := websocket.Config{
		HandshakeTimeout: 800 * time.Second,
		ReadBufferSize:   1824,
		WriteBufferSize:  256,
	}
	app.Static("/static", "./views/static")
	app.Static("/public", "./views/public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "VFtalk",
		})
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login/index", fiber.Map{
			"Title": "Login",
			"Desc":  "Welcome back, please enter your creds",
		})
	})
	app.Use("/room", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/room", websocket.New(HandleClients, wsConf))
	addr := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	log.Fatal(app.Listen(addr))
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("Error occured: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
