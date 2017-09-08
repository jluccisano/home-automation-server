package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
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
	subConfig := config.GetConf("test")
	loadKeys(subConfig)
	http.HandleFunc("/authenticate", GetTokenHandler)
	http.HandleFunc("/fakeUrl", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Listening...")
	go func() {
		fatal(http.ListenAndServe(":8515", nil))
	}()
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
	token, err := createToken(User{Username:"foo",Password:"bar"})

	var jsonStr = []byte(`{"username":"foo","password":"bar"}`)
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
	token, err := createToken(User{Username:"foo",Password:"bar"})
	req, err := http.NewRequest("GET","http://localhost:8515/fakeUrl", nil)
	fatal(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)
	assert.Equal(res.StatusCode, http.StatusOK, "status should be equal")
	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(body), "OK", "status should be equal")
}

func TestAuthMiddlewareWithNoHeader(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "/fakeUrl", nil)
	if err != nil {
		t.Fatal(err)
	}
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r)
		assert.Equal(w.Header().Get("Unauthorized"), "401", "Unauthorized");
	})
	rr := httptest.NewRecorder()
	handler := authMiddleware(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestAuthMiddlewareWithValidHeader(t *testing.T) {
	assert := assert.New(t)
	token, err := createToken(User{Username:"foo",Password:"bar"})
	req, err := http.NewRequest("GET", "/fakeUrl", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	if err != nil {
		t.Fatal(err)
	}
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r)
		assert.Equal(w.Header().Get("Unauthorized"), "401", "Unauthorized");
	})
	rr := httptest.NewRecorder()
	handler := authMiddleware(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}