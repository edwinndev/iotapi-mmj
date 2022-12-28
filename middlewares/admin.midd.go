package middlewares

import (
	"net/http"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/handlers"
)

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		claims, err := handlers.TokenValidate(w, token)
		if err != nil {
			return
		}
		if claims.Issuer == commons.RoleMaster {
			next.ServeHTTP(w, r)
		} else {
			commons.ApiNotFound(w, "Accesso denegado, permisos insuficientes")
		}
	}
}
