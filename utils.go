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
	PublicKey string `yaml:"publicKey"`
	PrivateKey string `yaml:"privateKey"`
}

type Config struct {
	Prod *SubConfig
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
