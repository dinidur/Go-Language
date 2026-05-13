package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string  `json:"name"`
	Subject string  `json:"subject"`
	Grade   float64 `json:"grade"`
}