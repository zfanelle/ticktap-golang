package service

import (
	"log"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	pb "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config/protos"
	transaction "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	repository "github.com/zfanelle/ticktap-golang/internal/app/mainservice/repository"
	"golang.org/x/net/context"
)

func CreateTransaction(appConfig config.AppConfig, newtransaction transaction.Transaction) {

	createTickets(appConfig, newtransaction)

	repository.CreateTransaction(appConfig, newtransaction)

}

func GetTransaction(appConfig config.AppConfig, transactionID int) transaction.Transaction {

	transaction := repository.GetTransaction(appConfig, transactionID)

	return transaction

}

func GetAllTransactions(appConfig config.AppConfig) []transaction.Transaction {

	transactions := repository.GetAllTransactions(appConfig)

	return transactions
}

func createTickets(appConfig config.AppConfig, newTransaction transaction.Transaction) {

	grpcClient := *appConfig.GrpcClient

	ticket := &pb.TicketCreationRequest{Account: int32(newTransaction.Account), Event: int32(newTransaction.Event), Quantity: int32(newTransaction.Quantity)}

	response, err := grpcClient.CreateTickets(context.Background(), ticket)

	if err != nil {
		log.Fatalf("Error when calling Ticketing Service: %s", err)
	}
	log.Printf("Response from server: %s", response.String())
}
