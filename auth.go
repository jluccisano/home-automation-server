package main


import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
)

func registerAuth(r *mux.Router) {
	r.Handle("/authenticate", GetTokenHandler).Methods("POST")
}

/* Set up a global string for our secret */
const (
	ValidUser = "John"
	ValidPass = "Doe"
	SecretKey = "WOW,MuchShibe,ToDogge"
)

type User struct{
	Id      string
	Password string
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
		/* Create the token */
		token := jwt.New(jwt.SigningMethodHS256)

		/* Set token claims */
		claims := make(jwt.MapClaims)
		claims["admin"] = true
		claims["name"] = "Ado Kukic"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		/* Sign the token with our secret */
		tokenString, err := token.SignedString(SecretKey)

		/* Finally, write the token to the browser window */

		if err != nil {
			fmt.Println(err)
			fmt.Println("User or Password is not valid:", token)
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

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		reqToken = splitToken[1]

		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			return []byte(SecretKey), nil
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

