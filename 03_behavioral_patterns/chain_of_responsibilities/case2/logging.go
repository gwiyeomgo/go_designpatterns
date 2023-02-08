package main

import "fmt"

type Logging struct {
	next Handler
}

func (a *Logging) execute(r *Request) {
	fmt.Println("로깅!")
	a.next.execute(r)
}

func (a *Logging) setNext(next Handler) {
	a.next = next
}
