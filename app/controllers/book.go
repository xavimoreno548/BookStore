package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/app/models"
	"github.com/xavimoreno548/BookStore/db"
)

type Book interface{
	GetBooks()
	CreateBook()
	FindBook()
	UpdateBook()
	DeleteBook()
}

func GetBooks(ctx *gin.Context)(){
	var books []models.Book
	DB := db.GetConnDB()

	DB.Find(&books)

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(ctx *gin.Context){
	// Create an instance of input for validations
	var input models.BookInputCreate

	// Data validation
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB := db.GetConnDB()

	book := models.Book{
		Name:      		input.Name,
		ISBN:    		input.ISBN,
		Price: 			input.Price,
		AuthorID:    	input.AuthorID,
		EditorialID: 	input.EditorialID,
	}

	DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(ctx *gin.Context){
	var book models.Book
	DB := db.GetConnDB()

	if err := DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context){
	var book models.Book
	DB := db.GetConnDB()

	if err := DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found, " + err.Error()})
		return
	}

	fmt.Printf("%v/n", book)

	var input models.BookInputUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Model(&book).Updates(book.ConvertModelToInterface(input)).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context){
	var book models.Book
	DB := db.GetConnDB()

	if err := DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}