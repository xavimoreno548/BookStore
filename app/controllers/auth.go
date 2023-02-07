package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/app/models"
	"github.com/xavimoreno548/BookStore/db"
	"github.com/xavimoreno548/BookStore/internal/services"
)

func Register(ctx *gin.Context){
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	dbConn := db.GetConnDB()
	_, err := user.Save(dbConn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "validated"})
}

func Login(ctx *gin.Context){
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Email = input.Email
	user.Password = input.Password

	dbConnection := db.GetConnDB()
	token, err := models.LoginCheck(user.Email, user.Password, dbConnection)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email or password is invalid."})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(ctx *gin.Context){
	userID, err := services.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConn := db.GetConnDB()
	user, err := models.GetUserByID(userID, dbConn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}