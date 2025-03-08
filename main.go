package main

import (
	"library-api/database"
	"library-api/routes"
	"fmt"
	"log"
	"os"
)

func main() {
	
	dbString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if dbString == "" || port == "" {
		log.Fatal("Required environment variables are not set.")
	}

	// Connect to database and migrate
	database.ConnectDb(dbString)
	database.Migrate()

	// Initialize and run the router
	router := routes.InitRouter()
	portNum := fmt.Sprintf(":%s", port)
	router.Run(portNum)
}
