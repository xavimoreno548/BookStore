package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name string `json:"name"`
	ISBN string `json:"isbn"`
	Price int `json:"price"`
	AuthorID int64 `json:"author_id"`
	EditorialID int64 `json:"editorial_id"`
}

type BookInputCreate struct {
	Name string `json:"name" binding:"required"`
	ISBN string `json:"isbn" binding:"required"`
	Price int `json:"price" binding:"required"`
	AuthorID int64 `json:"author_id" binding:"required"`
	EditorialID int64 `json:"editorial_id" binding:"required"`
}

type BookInputUpdate struct {
	Name string `json:"name"`
	ISBN string `json:"isbn"`
	AuthorID int64 `json:"author_id"`
	EditorialID int64 `json:"editorial_id"`
}

func (b Book) ConvertModelToInterface(input BookInputUpdate) map[string]interface{} {

	newInput := b

	if input.Name != "" {
		newInput.Name = input.Name
	}

	if input.ISBN != "" {
		newInput.ISBN = input.ISBN
	}

	if input.AuthorID != 0{
		newInput.AuthorID = input.AuthorID
	}

	if input.EditorialID != 0 {
		newInput.EditorialID = input.EditorialID
	}

	return map[string]interface{}{
		"name": newInput.Name,
		"isbn": newInput.ISBN,
		"author_id": newInput.AuthorID,
		"editorial_id": newInput.EditorialID,
	}
}