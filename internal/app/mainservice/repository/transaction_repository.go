package repsitory

import (
	"fmt"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	transaction "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
)

func CreateTransaction(appConfig config.AppConfig, newTransaction transaction.Transaction) {

	db := appConfig.DB

	sql := fmt.Sprintf("INSERT INTO ticktap.transaction (event,account) VALUES (%d, %d)", newTransaction.Event, newTransaction.Account)

	db.MustExec(sql)

}

func GetTransaction(appConfig config.AppConfig, TransactionID int) transaction.Transaction {
	db := appConfig.DB

	Transactions := []transaction.TransactionDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.transaction WHERE ID = %d", TransactionID)

	db.Select(&Transactions, sql)

	// Add in error checking

	Transaction := Transactions[0].DBtoTransaction()

	return Transaction
}

func GetAllTransactions(appConfig config.AppConfig) []transaction.Transaction {
	db := appConfig.DB

	dbTransactions := []transaction.TransactionDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.transaction")

	db.Select(&dbTransactions, sql)

	Transactions := []transaction.Transaction{}

	// Add in error checking

	for _, Transaction := range dbTransactions {
		Transactions = append(Transactions, Transaction.DBtoTransaction())
	}

	return Transactions
}
