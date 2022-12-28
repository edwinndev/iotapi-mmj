package middlewares

import (
	"net/http"

	"github.com/edwinndev/iotapi-mmj/handlers"
)

func JWTMiddlweare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := handlers.TokenValidate(w, token)
		if err != nil {
			return
		}
		next.ServeHTTP(w, r)
	}
}
