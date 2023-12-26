package configs

import (
	"os"
	"strconv"
)

type MailerConf struct {
	DefaultFromEmail string
	DefaultFromName  string
	ReplyToEmail     string
	UseBcc           bool
	DefaultMailer    string
}

const (
	MailerDockerMailserver = `dockermailserver`
	MailerMailhog          = `mailhog`
)

func EnvMailer() MailerConf {
	useBcc, _ := strconv.ParseBool(os.Getenv("MAILER_USE_BCC"))
	return MailerConf{
		DefaultFromEmail: os.Getenv("MAILER_DEFAULT_FROM_EMAIL"),
		DefaultFromName:  os.Getenv("MAILER_DEFAULT_FROM_NAME"),
		ReplyToEmail:     os.Getenv("MAILER_REPLY_TO_EMAIL"),
		UseBcc:           useBcc,
		DefaultMailer:    os.Getenv("MAILER_DEFAULT"),
	}

}
