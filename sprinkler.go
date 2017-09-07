package main

import (
	"fmt"
	"os/exec"
	"net/http"
	"github.com/gorilla/mux"
)

func registerSprinklerControl(router *mux.Router) {

	var ZonesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args := []string{"get"}
		cmd := exec.Command("/opt/relay_control/relay_control.py", args...)
		out,err := cmd.Output()
		if err != nil {
			println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(out))
	})

	var ZoneHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args := []string{"get"}
		vars := mux.Vars(r)
		id := vars["id"]
		if id != "" {
			args = append(args,"--relay")
			args = append(args,(fmt.Sprintf("%s", id)))
		}
		cmd := exec.Command("/opt/relay_control/relay_control.py", args...)
		out,err := cmd.Output()
		if err != nil {
			println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(out))
	})

	var SetZoneHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		vals := r.URL.Query()
		state, ok := vals["state"]
		if ok {
			println("State param is mandatory.")
		}

		args := []string{"set","--state", fmt.Sprintf("%s", state)}
		if !ok {
			args = append(args,"--relay")
			args = append(args, fmt.Sprintf("%s", id))
		}
		cmd := exec.Command("/opt/relay_control/relay_control.py", args...)
		_, err := cmd.Output()

		if err != nil {
			println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("OK"))
	})

	router.Handle("/sprinkler/zones", authMiddleware(ZonesHandler)).Methods("GET")
	router.Handle("/sprinkler/zones/{id}", authMiddleware(ZoneHandler)).Methods("GET")
	router.Handle("/sprinkler/zones/{id}", authMiddleware(SetZoneHandler)).Methods("POST")
}
