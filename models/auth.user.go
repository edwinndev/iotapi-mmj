package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Names    string `json:"names"`
	Surnames string `json:"surnames"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}
