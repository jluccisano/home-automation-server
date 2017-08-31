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

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./public/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))


	registerAuth(r, subConfig)
	registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", handlers.LoggingHandler(os.Stdout, r))

}
