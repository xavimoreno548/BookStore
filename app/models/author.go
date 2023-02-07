package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name"`
	Books []Book
}

type AuthorCreateInput struct {
	gorm.Model
	Name string `json:"name"`
}