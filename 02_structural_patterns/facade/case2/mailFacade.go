package case2

type MailFacade struct {
	settings    *Settings
	Message     *Message
	emailSender *EmailSender
}

func NewEmailSender(settings Settings) *EmailSender {
	return &EmailSender{settings: settings}
}
