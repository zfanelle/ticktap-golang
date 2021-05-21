package controller

import (
	"fmt"
	"net/http"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
)

func TestEndpoint(config config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Application is working!")
	}

}
