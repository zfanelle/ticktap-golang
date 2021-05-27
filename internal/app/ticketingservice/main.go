package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/zfanelle/ticktap-golang/internal/app/mainservice/protos"
	"google.golang.org/grpc"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedTicketingServiceServer
}

func (s *server) CreateTickets(ctx context.Context, in *pb.TicketCreationRequest) (*empty.Empty, error) {

	log.Printf("Received: Request")

	return &empty.Empty{}, nil

}

func (s *server) GetTicket(ctx context.Context, in *pb.TicketId) (*pb.Ticket, error) {

	log.Printf("Getting ticket")

	return &pb.Ticket{}, nil

}

func (s *server) GetAllTickets(ctx context.Context, none *empty.Empty) (*pb.Tickets, error) {

	log.Printf("Getting all tickets")

	return &pb.Tickets{}, nil

}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTicketingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
