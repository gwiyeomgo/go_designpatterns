package main

type Handler interface {
	execute(*Request)
	setNext(Handler)
}
