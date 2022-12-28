package handlers

import (
	"strconv"
	"time"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/models"
	"github.com/golang-jwt/jwt/v4"
)

func MakeToken(user models.User) (string, error) {
	var claims = &jwt.RegisteredClaims{
		ID:        strconv.Itoa(int(user.ID)),
		Subject:   user.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		Issuer:    user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(commons.PrivateKey))
}
