package configs

import (
	"os"
	"strconv"
)

type MailhogConf struct {
	MailerConf
	MailhogHost string
	MailhogPort int
}

func EnvMailhog() MailhogConf {
	mhPort, _ := strconv.Atoi(os.Getenv("MAILHOG_PORT"))
	return MailhogConf{
		MailhogHost: os.Getenv("MAILHOG_HOST"),
		MailhogPort: mhPort,
		MailerConf:  EnvMailer(),
	}
}
