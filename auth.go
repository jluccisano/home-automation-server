package main

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
	"io/ioutil"
	"crypto/rsa"
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

	r.Handle("/authenticate", addDefaultHeaders(GetTokenHandler)).Methods("POST")
}

/* Set up a global string for our secret */
const (
	ValidUser = "foo"
	ValidPass = "bar"
)

type User struct{
	Id      string
	Password string
}

func createToken (user User) (string, error) {
	/* Create the token */
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	/* Set token claims */
	claims := make(jwt.MapClaims)
	claims["admin"] = true
	claims["username"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
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

	if user.Id == ValidUser && user.Password == ValidPass {

		tokenString, err := createToken(user)
		/* Finally, write the token to the browser window */
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

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		reqToken = splitToken[1]

		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			return verifyKey, nil
		})

		if err == nil && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			fmt.Println("Token is not valid:")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	})
}

