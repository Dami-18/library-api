package middlewares

import (
	"library-api/auth"
	"library-api/database"
	"library-api/models"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
		  context.JSON(401, gin.H{"error": "request does not contain an access token"})
		  context.Abort()
		  return
		}
		claims, err := auth.ValidateTokenAndGetClaims(tokenString)
		if err != nil {
		  context.JSON(401, gin.H{"error": err.Error()})
		  context.Abort()
		  return
		}

		userId := claims.ID
		context.Set("userId",userId)
		context.Next()
	  }	
}

func AdminOnly() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, exists := context.Get("userId")

		if !exists {
			context.JSON(401, gin.H{"error":"user id not found in token claims!"})
		}

		var user models.User
		if err := database.Instance.First(&user,userID).Error; err != nil {
			context.JSON(401, gin.H{"error":"user not found"})
			context.Abort()
			return
		}
		if user.Role != "admin" {
			context.JSON(403, gin.H{"error":"admin access required!"})
			context.Abort()
			return
		}
		context.Next()
	}
}
