package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	var config Config
	subConfig := config.GetConf("prod")

	r := mux.NewRouter()

	registerAuth(r, subConfig)
	registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", handlers.LoggingHandler(os.Stdout, r))

}
