package main

import (
	"log"
	"runtime"

	"vftalk/conf"
	"vftalk/presentation"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cpu := runtime.NumCPU()
	log.Println("Total CPU Cores : ", cpu)
	runtime.GOMAXPROCS(cpu)
}

func main() {
	ws := &presentation.WebServer{
		AppName: "VFtalk - Chat App",
		Cfg:     conf.EnvWebConf(),
	}

	ws.Start()
}
