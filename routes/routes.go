package routes

import (
	"OpenSoft-MT/controllers"
	"OpenSoft-MT/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token",controllers.GenerateToken)
		api.POST("/user/register",controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping",controllers.Ping)
		}
	}
	return router
}