package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// function to check if server is running
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "ping-pong"})
}
