package main

import "fmt"

type CoinState struct {
	ticketMachine *TicketMachine
}

func (n CoinState) insertCoin() {
	fmt.Println("이미 동전이 들어있습니다.")
}

func (n CoinState) printTicket() {
	fmt.Println("프린트")
	n.ticketMachine.setState(n.ticketMachine.noCoinState)
}
