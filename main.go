package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	var config Config
	subConfig := config.GetConf("Prod")

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	registerAuth(r, subConfig)
	registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", handlers.LoggingHandler(os.Stdout, r))

}
