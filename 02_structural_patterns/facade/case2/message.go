package main

type Message struct {
	from string
	to   []string
	msg  []byte
}

func newMessage(from string, to []string, msg []byte) *Message {
	return &Message{
		from: from,
		to:   to,
		msg:  msg,
	}
}
