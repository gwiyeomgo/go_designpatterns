package main

import (
	"fmt"
	"net/smtp"
)

type MailFacade struct {
	settings    *Settings
	message     *Message
	emailSender *EmailSender
}

func newMailFacade(password string, from string, to []string, msg []byte) *MailFacade {
	return &MailFacade{
		settings:    newSettings(password),
		message:     newMessage(from, to, msg),
		emailSender: newEmailSender(),
	}
}

func (f *MailFacade) Send() {
	addr := f.settings.ServerHost + ":" + f.settings.ServerPort
	auth := smtp.PlainAuth("", f.message.from, f.settings.Password, f.settings.ServerHost)
	/*err := f.emailSender.send(addr, auth, f.message.from, f.message.to, f.message.msg)
	if err != nil {
		return err
	}*/
	result := fmt.Sprintf("addr:%s,auth:%s,from:%s,to:%s,msg:%s", addr, auth, f.message.from, f.message.to, f.message.msg)
	fmt.Println(result)
}
