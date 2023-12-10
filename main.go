package main

import (
	"vftalk/conf"
	"vftalk/presentation"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	var zlog zerolog.Logger = conf.GetLogger()
	err := godotenv.Load(".env")
	if err != nil {
		zlog.Panic().
			Str("ERROR", err.Error()).
			Msg("cannot load .env files")
	}

	ws := &presentation.WebServer{
		AppName: "VFtalk - Chat App",
		Cfg:     conf.EnvWebConf(),
	}

	ws.Start()
}
