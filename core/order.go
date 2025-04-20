package core

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Total float64 `json:"total"`
}
