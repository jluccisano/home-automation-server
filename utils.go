package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Url string `yaml:"url"`
	User string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	PrivateKey string `yaml:"privateKey"`
	PublicKey string `yaml:"publicKey"`
}

func (c *Config) GetConf(env string) *Config {
	yamlFile, err := ioutil.ReadFile("conf/conf-"+ env + ".yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
