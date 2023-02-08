package main

import "fmt"

type Print struct {
	next Handler
}

func (a *Print) execute(r *Request) {
	fmt.Println("요청내용:" + r.message)
}

func (a *Print) setNext(next Handler) {
	a.next = next
}
