package main

import (
	"log"
	"github.com/ungerik/go-rest"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
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

	registerAlarmControl(c)
	registerSprinklerControl()

	stopServerChan := make(chan struct{})

	fmt.Printf("Starting REST server\n")

	rest.HandleGET("/close", func() string {
		stopServerChan <- struct{}{}
		return "Stopping REST server..."
	})

	rest.RunServer("0.0.0.0:8080", stopServerChan)
}
