package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"reflect"
)

type SubConfig struct {
	Url string `yaml:"url"`
	User string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	PrivateKey string `yaml:"privateKey"`
	PublicKey string `yaml:"publicKey"`
}

type Config struct {
	Dev *SubConfig
	Test *SubConfig
}

func (c *Config) GetConf(env string) *SubConfig {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return reflect.ValueOf(c).Elem().FieldByName(env).Interface().(*SubConfig)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
