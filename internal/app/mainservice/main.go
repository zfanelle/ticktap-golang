package main

import (
	"log"
	"net/http"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	controller "github.com/zfanelle/ticktap-golang/internal/app/mainservice/controller"
)

func main() {

	appConfig := &config.AppConfig{}

	appConfig.Configure()

	appConfig.Router = controller.ConfigureRoutes(*appConfig)

	log.Fatal(http.ListenAndServe(":8080", appConfig.Router))

}
