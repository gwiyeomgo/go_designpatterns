package case2

import (
	"net/smtp"
)

type EmailSender struct {
	settings Settings
	send     func(string, smtp.Auth, string, []string, []byte) error
}

func (e *EmailSender) Send(to []string, body []byte) error {
	addr := e.settings.ServerHost + ":" + e.settings.ServerPort
	auth := smtp.PlainAuth("", e.settings.Username, e.settings.Password, e.settings.ServerHost)
	return e.send(addr, auth, e.settings.SenderAddr, to, body)
}
