package routes

import (
	"library-api/controllers"
	"library-api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/lib-api")
	{
		api.POST("/login",controllers.GenerateToken)
		api.POST("/signup",controllers.RegisterUser)

		api.POST("/books",middlewares.Auth(), middlewares.AdminOnly(), controllers.AddBook)
		api.GET("/books",controllers.GetBooks)
		api.GET("/books:id",controllers.GetBookByID) // need to specify optional paramters
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping",controllers.Ping)
		}
	}
	return router
}