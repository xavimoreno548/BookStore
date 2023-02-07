package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/app/models"
	"github.com/xavimoreno548/BookStore/db"
)

type Author interface {
	GetAuthors()
	CreateAuthor()
	FindAuthor()
	UpdateAuthor()
	DeleteAuthor()
}

func GetAuthors(ctx *gin.Context) {
	var authors []models.Author

	DB := db.GetConnDB()
	err := DB.Find(&authors).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": authors})
}

func CreateAuthor(ctx *gin.Context){
	var input models.AuthorCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB := db.GetConnDB()

	author := models.Author{
		Name: input.Name,
	}

	DB.Create(&author)

	ctx.JSON(http.StatusOK, gin.H{"data": author})
}