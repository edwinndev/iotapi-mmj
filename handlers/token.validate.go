package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/golang-jwt/jwt/v4"
)

func TokenValidate(w http.ResponseWriter, token string) (*jwt.RegisteredClaims, error) {
	tokenSplit := strings.Split(token, "Bearer")
	if len(tokenSplit) != 2 {
		commons.ApiUnauthorized(w, "Es necesario un token para acceder")
		return nil, errors.New("invalid token")
	}

	claims := &jwt.RegisteredClaims{}
	withClaims, err := jwt.ParseWithClaims(strings.TrimSpace(tokenSplit[1]), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(commons.PrivateKey), nil
	})
	if err != nil {
		commons.ApiUnauthorized(w, "Token invalido")
		return nil, errors.New("invalid token")
	}

	exists := database.UserExists(claims.Subject)
	if exists != nil {
		commons.ApiNotFound(w, "El sujeto del token es invalido")
		return nil, errors.New("invalid token")
	}

	if !withClaims.Valid {
		commons.ApiUnauthorized(w, "Acceso denegado")
		return nil, errors.New("access denied")
	}
	return claims, nil
}
