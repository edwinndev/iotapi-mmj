package models

type Device struct {
	Model
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
