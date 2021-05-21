package service

import (
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	transaction "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	repository "github.com/zfanelle/ticktap-golang/internal/app/mainservice/repository"
)

func CreateTransaction(appConfig config.AppConfig, newtransaction transaction.Transaction) {

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
