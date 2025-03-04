package routes

import (
	"library-api/controllers"
	"library-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// configure cors 
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	api := router.Group("/lib-api")
	{
		api.POST("/login",controllers.GenerateToken)
		api.POST("/signup",controllers.RegisterUser)

		api.POST("/books",middlewares.Auth(), middlewares.AdminOnly(), controllers.AddBook)
		api.GET("/books",controllers.GetBooks)
		api.GET("/books/:id",controllers.GetBookByID) // need to specify optional paramters after another '/:'

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping",controllers.Ping)
		}
	}
	return router
}