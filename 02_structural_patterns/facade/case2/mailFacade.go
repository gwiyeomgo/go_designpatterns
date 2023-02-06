package main

import (
	"fmt"
	"net/smtp"
)

// 파사드 하위 시스템 기능의 특정 부분에 대한 편리한 접근을 제공
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

func (f *MailFacade) Send() error {
	addr := f.settings.ServerHost + ":" + f.settings.ServerPort
	auth := smtp.PlainAuth("", f.message.from, f.settings.Password, f.settings.ServerHost)
	err := f.emailSender.send(addr, auth, f.message.from, f.message.to, f.message.msg)
	if err != nil {
		return err
	}
	result := fmt.Sprintf("addr:%s,auth:%s,from:%s,to:%s,msg:%s", addr, auth, f.message.from, f.message.to, f.message.msg)
	fmt.Println(result)
	return nil
}
