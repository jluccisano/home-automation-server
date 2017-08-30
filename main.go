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
	http.Handle("/", &MyServer{r})
	//r.Handle("/", http.FileServer(http.Dir("./views/")))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	registerAuth(r, subConfig)
	registerAlarmControl(r, subConfig)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", handlers.LoggingHandler(os.Stdout, r))

}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}