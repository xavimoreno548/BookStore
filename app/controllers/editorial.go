package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/app/models"
	"github.com/xavimoreno548/BookStore/db"
)

type Editorial interface {
	GetEditorials()
	CreateEditorial()
	FindEditorial()
	UpdateEditorial()
	DeleteEditorial()
}

func GetEditorials(ctx *gin.Context) {
	var editorials []models.Editorial
	DB := db.GetConnDB()

	DB.Find(&editorials)

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": editorials})
}

func CreateEditorial(ctx *gin.Context){
	var input models.EditorialInputCreate

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB := db.GetConnDB()

	editorial := models.Editorial{
		Name: input.Name,
	}

	DB.Create(&editorial)

	ctx.JSON(http.StatusOK, gin.H{"data": editorial})
}