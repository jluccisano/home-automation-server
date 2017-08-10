package main

import (
	"testing"
	"net/http"
	"log"
	"fmt"
	"net/url"
	"bytes"
	"io/ioutil"
)

func TestMain(m *testing.M) {
	http.ListenAndServe(":8515", nil)
	log.Println("Listening...")
}

func TestUnauthorizedIfNoToken(t *testing.T) {
	req, err := http.NewRequest("GET","http://localhost:8515/sprinkler/zones", nil)
	fatal(err)
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}

func TestAuthenticate(t *testing.T) {
	var jsonStr = []byte(`{"user":"foo","password":"bar"}`)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%v/authenticate", "8515"), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	fatal(err)
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}

func TestAuthorizedIfToken(t *testing.T) {
	req, err := http.NewRequest("GET","http://localhost:8515/sprinkler/zones", nil)
	fatal(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}


func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
