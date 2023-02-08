package main

import "fmt"

type Auth struct {
	next Handler
}

func (a *Auth) execute(r *Request) {
	fmt.Println("인증이 되었는가?")
	a.next.execute(r)
}

func (a *Auth) setNext(next Handler) {
	a.next = next
}
