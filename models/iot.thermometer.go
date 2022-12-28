package models

type Thermometer struct {
	Model
	Value  float32 `json:"value"`
	Device string  `json:"device"`
}
