package models

import "gorm.io/gorm"

type Phase struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
