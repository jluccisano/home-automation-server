package main

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

type conf struct {
	Url string `yaml:"url"`
	User string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	var c conf
	c.getConf()

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	registerAuth(r)
	registerAlarmControl(r, c)
	registerSprinklerControl(r)

	http.ListenAndServe(":8514", handlers.LoggingHandler(os.Stdout, r))

}
