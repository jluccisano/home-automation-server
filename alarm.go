package home_automation_server

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"github.com/ungerik/go-rest"
	"github.com/jluccisano/syno-cli/synoapi"
)

func AlarmController(c conf) {

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
