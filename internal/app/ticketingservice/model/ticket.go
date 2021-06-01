package model

import (
	pb "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/protos"
)

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

func (ticket Ticket) ProtoCreateTicketToTicket(ticketCreation *pb.TicketCreationRequest) Ticket {

	return Ticket{Account: int(ticketCreation.GetAccount()), Event: int(ticketCreation.GetEvent()), Quantity: int(ticketCreation.GetQuantity())}
}

func TicketToProtoTicket(ticket Ticket) pb.Ticket {

	return pb.Ticket{Id: int32(ticket.Id), Event: int32(ticket.Event), Account: int32(ticket.Account)}
}
