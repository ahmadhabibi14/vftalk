package mailer

import (
	"log"
	"vftalk/configs"

	"github.com/rs/zerolog"
	"github.com/wneessen/go-mail"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func NewMailer(l *zerolog.Logger) Mailer {
	mlr := Mailer{}
	mh, err := NewMailhog(configs.EnvMailhog())
	if err != nil {
		l.Error().Str(`Error: `, err.Error()).Msg(`Cannot load mailhog`)
	}
	mlr.SendMailFunc = mh.SendEmail
	return mlr
}

func (m *Mailer) SendUserRegisterEmail(email string) error {
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Welcome to VFtalk – Registration Successful!`,
		`Halooo tess`,
		`<p>Halooo</p>`,
	)
}

func SendUserRegisterMail() {
	m := mail.NewMsg()
	if err := m.From("ahmadhabibi04@proton.me"); err != nil {
		log.Fatalf("Failed to set From address: %s", err)
	}
	if err := m.To("habi@ternaklinux.com"); err != nil {
		log.Fatalf("Failed to set To address: %s", err)
	}
	m.Subject("This is my first mail with go-mail!")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
	c, err := mail.NewClient(
		"127.0.0.1",
		mail.WithPort(1025),
		mail.WithTLSPolicy(mail.NoTLS),
	)
	if err != nil {
		log.Fatalf("Failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send mail: %s", err)
	}
}
