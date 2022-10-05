package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venkat7668/Go-JWT/database"
	"github.com/venkat7668/Go-JWT/models"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error0": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		context.Abort()
		return
	}

	db := database.Instance.Create(&user)
	if db.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error2": db.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
