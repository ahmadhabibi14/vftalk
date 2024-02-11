package mailer

import (
	"os"

	"github.com/rs/zerolog"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func NewMailer(l *zerolog.Logger) Mailer {
	var mlr Mailer = Mailer{}
	if os.Getenv("WEB_ENV") == "dev" {
		mh, err := NewMailhog()
		if err != nil {
			l.Error().Str("error", err.Error()).Msg(`cannot load mailhog`)
		} else {
			mlr.SendMailFunc = mh.SendEmail
		}
	} else {
		dms, err := NewDockermailserver()
		if err != nil {
			l.Error().Str("error", err.Error()).Msg(`cannot connect to dockermailserver`)
		} else {
			mlr.SendMailFunc = dms.SendEmail
		}
	}
	return mlr
}

func (m *Mailer) SendUserRegisterEmail(email string) error {
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Welcome to VFtalk`,
		`Registration Successful!`,
		`<p>Registration Successful!</p>`,
	)
}
