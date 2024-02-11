package configs

import (
	"os"
	"strconv"
)

type DockermailserverConf struct {
	MailerConf
	DockermailserverHost string
	DockermailserverPort int
	DockermailserverUser string
	DockermailserverPass string
}

func EnvDockermailserver() DockermailserverConf {
	port, _ := strconv.Atoi(os.Getenv("DOCKERMAILSERVER_PORT"))
	return DockermailserverConf{
		DockermailserverHost: os.Getenv("DOCKERMAILSERVER_HOST"),
		DockermailserverPort: port,
		DockermailserverUser: os.Getenv("DOCKERMAILSERVER_USER"),
		DockermailserverPass: os.Getenv("DOCKERMAILSERVER_PASS"),
		MailerConf:           EnvMailer(),
	}
}
