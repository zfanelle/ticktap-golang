package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	ticketService "github.com/zfanelle/ticktap-golang/internal/app/mainservice/service"
)

func GetTicket(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		string_id := params["id"]

		id, err := strconv.Atoi(string_id)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		ticket := ticketService.GetTicket(config, id)

		json.NewEncoder(w).Encode(ticket)
	}

}

func GetAllTickets(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tickets := ticketService.GetAllTickets(config)

		json.NewEncoder(w).Encode(tickets)
	}

}
