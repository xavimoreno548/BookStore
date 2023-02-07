package models

import "gorm.io/gorm"

type FavoriteItems struct {
	gorm.Model
	FavoriteID int64 `json:"favorite_id"`
	BookID int64 `json:"book_id"`
}
