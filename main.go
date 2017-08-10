package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	"log"
)

func main() {
	var config Config
	subConfig := config.GetConf("Prod")

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/home", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})).Methods("GET")

	registerAuth(r, subConfig)
	//registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	log.Println("Start Listening on 8514 over SSL...")

	go fatal(http.ListenAndServeTLS(":8514", "testResources/cert.pem", "testResources/key.pem", handlers.LoggingHandler(os.Stdout, r)))

	log.Println("Start Listening on 8515...")

	fatal(http.ListenAndServe(":8515", handlers.LoggingHandler(os.Stdout, r)))
	//http://www.kaihag.com/https-and-go/

}
