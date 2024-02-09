package main

import (
	"os"
	"strings"
	"vftalk/configs"
	_ "vftalk/docs"
	"vftalk/models"
	"vftalk/web"
)

// @title VFtalk API Docs
// @version 1.0
// @description Restful API Specification of VFtlak
// @termsOfService https://vftalk.my.id
// @contact.name API Support
// @contact.email habi@ternaklinux.com
// @license.name MIT License
// @license.url https://www.mit.edu/~amini/LICENSE.md
// @host vftalk.my.id
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
