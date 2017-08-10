package main

import (
	"testing"
	"net/http"
	"log"
	"fmt"
	"bytes"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)
var (
	serverPort int
	fakeUser = map[string]string{
		"foo": "bar",
	}
)

func init() {
	var config Config
	subConfig := config.GetConf("Test")
	loadKeys(subConfig)
	http.HandleFunc("/authenticate", GetTokenHandler)
	http.HandleFunc("/fakeUrl", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/testSSL", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK SSL"))
	})

	log.Println("Listening...")
	go func() {
		fatal(http.ListenAndServe(":8515", nil))
		fatal(http.ListenAndServeTLS(":8514", "testResources/cert.pem", "testResources/key.pem", nil))
		//openssl req -new -x509 -sha256 -key testResources/rs256-4096-private.rsa -out testResources/server.crt -days 3650
		//openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
		//openssl req -new -key <private key file name>.key -out <csr file name>.crt


	}()
}

func TestTLS(t *testing.T) {
	req, err := http.NewRequest("GET","https://localhost:8514/testSSL", nil)
	fatal(err)
	resp, err := http.DefaultClient.Do(req)
	fatal(err)
	fmt.Print(resp)
}

func TestUnauthorizedIfNoToken(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET","http://localhost:8515/fakeUrl", nil)
	fatal(err)
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	assert.Equal(res.StatusCode, http.StatusOK, "status should be equal")
}

func TestAuthenticate(t *testing.T) {
	assert := assert.New(t)
	token, err := createToken(User{Id:"foo",Password:"bar"})

	var jsonStr = []byte(`{"id":"foo","password":"bar"}`)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%v/authenticate", "8515"), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	fatal(err)
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	assert.Equal(res.StatusCode, http.StatusOK, "status should be equal")
	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(body), token, "token should be equal")

}

func TestAuthorizedIfToken(t *testing.T) {
	assert := assert.New(t)
	token, err := createToken(User{Id:"foo",Password:"bar"})
	req, err := http.NewRequest("GET","http://localhost:8515/fakeUrl", nil)
	fatal(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	assert.Equal(res.StatusCode, http.StatusOK, "status should be equal")
	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(body), "OK", "status should be equal")

}