package model

type Ticket struct {
	Id       int `json:"id"`
	Account  int `json:"account"`
	Event    int `json:"event"`
	Quantity int `json:"entity_type"`
}

type TicketDB struct {
	Id      int `db:"id"`
	Account int `db:"account"`
	Event   int `db:"event"`
}

func (ticketDB TicketDB) DBtoTicket() Ticket {
	return Ticket{Id: ticketDB.Id, Account: ticketDB.Account, Event: ticketDB.Event}
}
