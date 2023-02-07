package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID int64 `json:"user_id"`
	FavoriteItems []FavoriteItems
}
