package models

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID        uint      `json:"id"`
	Names     string    `json:"names"`
	Surnames  string    `json:"surnames"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`
	Token     string    `json:"token"`
}
