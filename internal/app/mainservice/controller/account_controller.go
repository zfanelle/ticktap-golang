package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	account "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	accountService "github.com/zfanelle/ticktap-golang/internal/app/mainservice/service"
)

func CreateAccount(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var account account.AccountDTO

		_ = json.NewDecoder(r.Body).Decode(&account)

		new_account := account.DTOtoAccount()

		accountService.CreateAccount(config, new_account)
	}

}

func GetAccount(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		string_id := params["id"]

		id, err := strconv.Atoi(string_id)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		account := accountService.GetAccount(config, id)

		json.NewEncoder(w).Encode(account)
	}

}

func GetAllAccounts(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		accounts := accountService.GetAllAccounts(config)

		json.NewEncoder(w).Encode(accounts)
	}

}
