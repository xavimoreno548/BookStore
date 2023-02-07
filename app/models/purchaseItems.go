package models

import "gorm.io/gorm"

type PurchaseItems struct {
	gorm.Model
	PurchaseID int64 `json:"purchase_id"`
	BookID int64 `json:"book_id"`
}
