package main

import (
	"log"
	"net/url"
	"github.com/jluccisano/go-rest"
	"github.com/jluccisano/syno-cli/synoapi"
)

func registerAlarmControl(c conf) {

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
			return err.Error();
		}
		return "OK";
	})

	rest.HandleGET("/disable", func(in url.Values)  string {
		err := client.Disable()
		if err != nil {
			log.Fatalf("Disable failed: %v", err)
			return err.Error();
		}
		return "OK";
	})
}
