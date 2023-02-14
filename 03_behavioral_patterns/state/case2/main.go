package main

func main() {
	ticketMachine := newTicketMachine()
	ticketMachine.printTicket()
	ticketMachine.insertCoin()
	ticketMachine.insertCoin()
	ticketMachine.printTicket()
}
