package main

import (
	"log"
	"github.com/jluccisano/syno-cli/synoapi"
	"net/http"
	"github.com/gorilla/mux"
)

func registerAlarmControl(r *mux.Router, c conf) {

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

	r.Handle("/alarm/enable", EnableHandler).Methods("GET")
	r.Handle("/alarm/disable", DisableHandler).Methods("GET")
}
