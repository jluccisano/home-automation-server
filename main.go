package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	"github.com/rs/cors"
)

func main() {
	var config Config
	subConfig := config.GetConf("prod")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	r := mux.NewRouter()

	registerAuth(r, subConfig)
	registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", c.Handler(handlers.LoggingHandler(os.Stdout, r)))

}
