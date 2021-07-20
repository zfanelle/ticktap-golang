package model

import (
	pb "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config/protos"
)

type Ticket struct {
	Id      int `json:"id"`
	Account int `json:"account"`
	Event   int `json:"event"`
}

func ProtoToTicket(protoTicket *pb.Ticket) Ticket {

	ticket := Ticket{Id: int(protoTicket.Id), Account: int(protoTicket.Account), Event: int(protoTicket.Event)}
	return ticket
}
