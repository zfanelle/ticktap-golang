package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	event "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	eventService "github.com/zfanelle/ticktap-golang/internal/app/mainservice/service"
)

func CreateEvent(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var event event.EventDTO

		_ = json.NewDecoder(r.Body).Decode(&event)

		new_event := event.DTOtoevent()

		eventService.Createevent(config, new_event)
	}

}

func GetEvent(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		string_id := params["id"]

		id, err := strconv.Atoi(string_id)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		event := eventService.Getevent(config, id)

		json.NewEncoder(w).Encode(event)
	}

}

func GetAllEvents(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		events := eventService.GetAllEvents(config)

		json.NewEncoder(w).Encode(events)
	}

}
