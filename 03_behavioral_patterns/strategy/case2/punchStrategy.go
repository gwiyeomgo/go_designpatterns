package main

import "fmt"

// concrete strategy
type PunchStrategy struct {
}

func (a *PunchStrategy) attack() {
	fmt.Println("I have strong punch")
}
