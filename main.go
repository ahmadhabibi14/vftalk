package main

import (
	"os"
	"strings"
	"vftalk/app"
	"vftalk/conf"
	"vftalk/models"
)

func main() {
	conf.LoadEnv()
	zlog := conf.InitLogger()
	validArgs := `web, migrate`

	var mode string
	if len(os.Args) < 2 {
		mode = `web`
	} else {
		mode = strings.ToLower(os.Args[1])
	}

	switch mode {
	case `web`:
		ws := &app.WebServer{
			AppName: "VFtalk - Chat App",
			Cfg:     conf.EnvWebConf(),
		}
		ws.Start()
	case `migrate`:
		models.RunMigration()
	default:
		zlog.Fatal().Msg(`Must start with: ` + validArgs)
	}
}
