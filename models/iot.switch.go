package models

type Switch struct {
	Model
	Device string `json:"device"`
	Value  bool   `json:"value"`
}
