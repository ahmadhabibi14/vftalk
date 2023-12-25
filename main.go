package main

import (
	"os"
	"strings"
	"vftalk/app"
	"vftalk/conf"
	"vftalk/handlers"
	"vftalk/models"
	"vftalk/models/mailer"
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

	var mlr mailer.Mailer
	mh, err := mailer.NewMailhog(conf.EnvMailhog())
	if err != nil {
		zlog.Error().Str(`Error: `, err.Error()).Msg(`Cannot load mailhog`)
	}

	mlr.SendMailFunc = mh.SendEmail
	h := handlers.Handler{
		Mailer: mailer.Mailer{
			SendMailFunc: mlr.SendMailFunc,
		},
		Log: zlog,
	}

	switch mode {
	case `web`:
		ws := &app.WebServer{
			Handler: h,
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
