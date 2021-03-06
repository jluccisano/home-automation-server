package main

import (
	"log"
	"github.com/jluccisano/syno-cli/synoapi"
	"net/http"
	"github.com/gorilla/mux"
)

func registerAlarmControl(r *mux.Router, c *Config) {

	client := synoapi.NewClient(c.Url)
	err2 := client.Login(c.User, c.Passwd)
	if err2 != nil {
		log.Fatal(err2)
	}

	var EnableHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := client.Enable()
		if err != nil {
			log.Fatalf("Enable failed: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("OK"))
	})

	var DisableHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := client.Disable()
		if err != nil {
			log.Fatalf("Enable failed: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("OK"))
	})

	var GetStatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
		err := client.Disable()
		if err != nil {
			log.Fatalf("Enable failed: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("OK"))
	})

	r.Handle("/alarm/enable", authMiddleware(EnableHandler)).Methods("GET")
	r.Handle("/alarm/disable", authMiddleware(DisableHandler)).Methods("GET")
	r.Handle("/alarm/status", authMiddleware(GetStatusHandler)).Methods("GET")
}
