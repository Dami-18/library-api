package controllers

import (
	"library-api/auth"
	"library-api/database"
	"library-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenRequest struct { // this are the payloads of request
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	// ID       uint   `json:"id" gorm:"primaryKey"` // primary key means on basis of this, book will be found in database, if not done explicit field querying
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil { // in should bind json, we pass the reference of that struct which will be sent as JSON data body via POST request
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if user already exists, if exists then check if password matches
	record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPasswd(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.ID) // if no such user, then register that user
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
