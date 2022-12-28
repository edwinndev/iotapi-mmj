package database

import (
	"github.com/edwinndev/iotapi-mmj/models"
)

func UserExists(email string) (exists error) {
	return Mysql.First(&models.User{}, "email=?", email).Error
}
