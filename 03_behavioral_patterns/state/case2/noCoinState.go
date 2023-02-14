package main

import "fmt"

type NoCoinState struct {
	ticketMachine *TicketMachine
}

func (n NoCoinState) insertCoin() {
	n.ticketMachine.setState(n.ticketMachine.coinState)
	fmt.Println("동전이 추가되었습니다.")
}

func (n NoCoinState) printTicket() {
	fmt.Println("동전이 없습니다. 동전을 넣어주세요")
}
