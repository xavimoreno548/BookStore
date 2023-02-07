package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	UserId int64 `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Code string `json:"code"`
	PurchaseItems []PurchaseItems
}
