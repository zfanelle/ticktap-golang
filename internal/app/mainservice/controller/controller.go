package controller

import (
	"github.com/gorilla/mux"
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
)

func ConfigureRoutes(appConfig config.AppConfig) *mux.Router {

	router := mux.NewRouter()
	router.Methods("GET").Path("/test").HandlerFunc(TestEndpoint(appConfig))
	router.Methods("POST").Path("/account").HandlerFunc(CreateAccount(appConfig))
	router.Methods("GET").Path("/account").HandlerFunc(GetAllAccounts(appConfig))
	router.Methods("GET").Path("/account/{id}").HandlerFunc(GetAccount(appConfig))
	router.Methods("POST").Path("/event").HandlerFunc(CreateEvent(appConfig))
	router.Methods("GET").Path("/event").HandlerFunc(GetAllEvents(appConfig))
	router.Methods("GET").Path("/event/{id}").HandlerFunc(GetAllEvents(appConfig))
	router.Methods("POST").Path("/transaction").HandlerFunc(CreateTransaction(appConfig))
	router.Methods("GET").Path("/transaction").HandlerFunc(GetAllTransactions(appConfig))
	router.Methods("GET").Path("/transaction/{id}").HandlerFunc(GetTransaction(appConfig))
	return router

}
