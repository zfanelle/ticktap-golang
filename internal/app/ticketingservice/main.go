package main

import (
	"context"
	"fmt"
	"log"
	"net"

	config "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/config"
	model "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/model"
	pb "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/protos"
	service "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/service"
	"google.golang.org/grpc"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedTicketingServiceServer
	AppConfig *config.AppConfig
}

func (s *server) CreateTickets(ctx context.Context, in *pb.TicketCreationRequest) (*empty.Empty, error) {

	ticket := model.Ticket{}

	ticket.ProtoCreateTicketToTicket(in)

	service.CreateTicket(*s.AppConfig, ticket)

	log.Printf("Creating Tickets")

	return &empty.Empty{}, nil

}

func (s *server) GetTicket(ctx context.Context, in *pb.TicketId) (*pb.Ticket, error) {

	log.Printf("Getting ticket")

	ticket := service.GetTicket(*s.AppConfig, int(in.GetId()))

	return &pb.Ticket{Id: int32(ticket.Id), Event: int32(ticket.Event), Account: int32(ticket.Account)}, nil

}

func (s *server) GetAllTickets(ctx context.Context, none *empty.Empty) (*pb.Tickets, error) {

	log.Printf("Getting all tickets")

	tickets := service.GetAllTickets(*s.AppConfig)

	protoTickets := pb.Tickets{}

	protoTicketsList := protoTickets.GetTickets()

	for _, s := range tickets {
		newTicket := model.TicketToProtoTicket(s)
		protoTicketsList = append(protoTicketsList, &newTicket)
	}

	protoTickets.Tickets = protoTicketsList

	return &protoTickets, nil

}

func main() {

	appConfig := &config.AppConfig{}

	appConfig.Configure()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTicketingServiceServer(grpcServer, &server{AppConfig: appConfig})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
