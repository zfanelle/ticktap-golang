package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	transaction "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	transactionService "github.com/zfanelle/ticktap-golang/internal/app/mainservice/service"
)

func CreateTransaction(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var transaction transaction.TransactionDTO

		_ = json.NewDecoder(r.Body).Decode(&transaction)

		new_transaction := transaction.DTOtoTransaction()

		transactionService.CreateTransaction(config, new_transaction)
	}

}

func GetTransaction(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		string_id := params["id"]

		id, err := strconv.Atoi(string_id)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		transaction := transactionService.GetTransaction(config, id)

		json.NewEncoder(w).Encode(transaction)
	}

}

func GetAllTransactions(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		transactions := transactionService.GetAllTransactions(config)

		json.NewEncoder(w).Encode(transactions)
	}

}
