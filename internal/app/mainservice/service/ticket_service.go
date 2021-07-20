package service

import (
	"log"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	pb "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config/protos"
	ticket "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	"golang.org/x/net/context"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

func GetTicket(appConfig config.AppConfig, ticketId int) ticket.Ticket {

	grpcClient := *appConfig.GrpcClient

	response, err := grpcClient.GetTicket(context.Background(), &pb.TicketId{Id: int32(ticketId)})

	if err != nil {
		log.Fatalf("Error when calling Ticketing Service: %s", err)
	}
	log.Printf("Response from server: %s", response.String())

	ticket := ticket.ProtoToTicket(response)

	return ticket

}

func GetAllTickets(appConfig config.AppConfig) []ticket.Ticket {

	grpcClient := *appConfig.GrpcClient

	response, err := grpcClient.GetAllTickets(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Error when calling Ticketing Service: %s", err)
	}
	log.Printf("Response from server: %s", response.String())

	var tickets []ticket.Ticket

	for _, pbTicket := range response.GetTickets() {
		tickets = append(tickets, ticket.ProtoToTicket(pbTicket))
	}

	return tickets
}
