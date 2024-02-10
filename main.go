package main

import (
	"os"
	"strings"
	"vftalk/configs"
	"vftalk/models"
	"vftalk/web"
)

func main() {
	configs.LoadEnv()
	zlog := configs.InitLogger()
	validArgs := `web, migrate-up, migrate-down`

	var mode string
	if len(os.Args) < 2 {
		mode = `web`
	} else {
		mode = strings.ToLower(os.Args[1])
	}

	switch mode {
	case `web`:
		ws := web.NewWebServer(configs.EnvWebConf(), zlog)
		ws.Start()
	case `migrate-up`:
		models.RunMigrationUp()
	case `migrate-down`:
		models.RunMigrationDown()
	default:
		zlog.Fatal().Msg(`Must start with: ` + validArgs)
	}
}
