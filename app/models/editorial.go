package models

import "gorm.io/gorm"

type Editorial struct {
	gorm.Model
	Name string `json:"name"`
	Books []Book
}

type EditorialInputCreate struct {
	gorm.Model
	Name string `json:"name"`
}