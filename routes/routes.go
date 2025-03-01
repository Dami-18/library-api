package routes

import (
	"library-api/controllers"
	"library-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/webstradev/gin-pagination/v2/pkg/pagination"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/lib-api")
	{
		api.POST("/login",controllers.GenerateToken)
		api.POST("/signup",controllers.RegisterUser)

		// to use gin-pagination middleware in api.GET() for pagination
		api.POST("/books",middlewares.Auth(), middlewares.AdminOnly(), controllers.AddBook)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping",controllers.Ping)
		}
	}
	return router
}