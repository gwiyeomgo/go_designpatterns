package main

type TicketMachine struct {
	noCoinState  State
	coinState    State
	currentState State
}

func (t *TicketMachine) setState(state State) {
	t.currentState = state
}

func (t *TicketMachine) printTicket() {
	t.currentState.printTicket()
}
func (t *TicketMachine) insertCoin() {
	t.currentState.insertCoin()
}
func newTicketMachine() *TicketMachine {
	t := &TicketMachine{}
	noCoinState := &NoCoinState{
		ticketMachine: t,
	}
	coinState := &CoinState{
		ticketMachine: t,
	}

	t.setState(noCoinState)
	t.noCoinState = noCoinState
	t.coinState = coinState
	return t
}
