package main

import (
	"github.com/jluccisano/syno-cli/synoapi"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"github.com/ungerik/go-rest"
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

	client := synoapi.NewClient(c.Url)
	err2 := client.Login(c.User, c.Passwd)
	if err2 != nil {
		log.Fatal(err2)
	}


	rest.HandleGET("/enable", func(in url.Values)  string {
		err := client.Enable()
		if err != nil {
			log.Fatalf("Enable failed: %v", err)
		}
		return err.Error();
	})

	rest.HandleGET("/disable", func(in url.Values)  string {
		err := client.Disable()
		if err != nil {
			log.Fatalf("Disable failed: %v", err)
		}
		return err.Error();
	})
}
