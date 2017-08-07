package home_automation_server

import (
	"log"
	"github.com/ungerik/go-rest"
	"net/url"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"home_automation_server/alarm"
	"home_automation_server/sprinkler"
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

	alarm := AlarmController(c)
	sprinkler := SprinklerController()

	// See RunServer below
	stopServerChan := make(chan struct{})

	fmt.Printf("Starting REST server\n")

	rest.HandleGET("/close", func() string {
		stopServerChan <- struct{}{}
		return "Stopping REST server..."
	})

	rest.RunServer("0.0.0.0:8080", stopServerChan)

}
