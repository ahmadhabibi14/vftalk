package main

import (
	"log"

	"vftalk/conf"
	"vftalk/presentation"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ws := &presentation.WebServer{
		AppName: "VFtalk - Chat App",
		Cfg:     conf.EnvWebConf(),
	}

	ws.Start()
}
