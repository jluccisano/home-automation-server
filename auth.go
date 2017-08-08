package main


import (
	"fmt"
	"net/http"
	"github.com/auth0-community/auth0"
	"gopkg.in/square/go-jose.v2"
)


func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte("{YOUR-AUTH0-API-SECRET}")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := "{YOUR-AUTH0-API-AUDIENCE}"

		configuration := auth0.NewConfiguration(secretProvider, []string{audience}, "https://{YOUR-AUTH0-DOMAIN}.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

