package case2

import "net/smtp"

type Message struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}
