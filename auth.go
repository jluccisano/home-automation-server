package main

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
)

func loadKeys(subConfig *Config) {

	signBytes, err := ioutil.ReadFile(subConfig.PrivateKey)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(subConfig.PublicKey)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func registerAuth(r *mux.Router, subConfig *Config) {
	loadKeys(subConfig)

	r.Handle("/authenticate", GetTokenHandler).Methods("POST")
}

const (
	ValidUser = "foo"
	ValidPass = "bar"
)

type User struct{
	Username      string
	Password string
}

func createToken (user User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := make(jwt.MapClaims)
	claims["admin"] = true
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(signKey)
}

/* Handlers */
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var user User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if user.Username == ValidUser && user.Password == ValidPass {
		tokenString, err := createToken(user)
		if err != nil {
			fmt.Println("User or Password is not valid:", tokenString)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			w.Write([]byte(tokenString))
		}
	}  else {
		fmt.Println(err)
		fmt.Println("User or Password is not valid:")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}
})


func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})
		if err == nil && reqToken.Valid {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			fmt.Fprint(w, "Unauthorized access, token is not valid")
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

