package main

import "net/smtp"

type EmailSender struct {
	send func(string, smtp.Auth, string, []string, []byte) error
}

func newEmailSender() *EmailSender {
	return &EmailSender{send: smtp.SendMail}
}
