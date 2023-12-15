package main

import (
	"os"
	"strings"
	"vftalk/conf"
	"vftalk/models"
	"vftalk/presentation"
)

func main() {
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
		ws := &presentation.WebServer{
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
