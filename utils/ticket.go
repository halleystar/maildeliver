package utils

type Ticket struct {
	total  int
	signal chan byte
}

func NewTicket() *Ticket {
	ticket := &Ticket{}
	ticket.total = Cfg.TicketMax
	ticket.signal = make(chan byte, ticket.total)
	for i := 0; i < ticket.total; i++ {
		ticket.signal <- 1
	}
}

func (ticket *Ticket) Add() {
	ticket.signal <- 1
}

func (ticket *Ticket) Done() {
	<-ticket.signal
}

func (ticket *Ticket) getTotal() int {
	return ticket.total
}

func (ticket *Ticket) getLeft() int {
	return len(ticket.signal)
}
