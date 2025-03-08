package main

import (
	"library-api/database"
	"library-api/routes"
	"fmt"
	
	"os"
)

func main() {

	dbUser := os.Getenv("DB_USER")
 	dbPassword := os.Getenv("DB_PASSWORD")
 	dbHost := os.Getenv("DB_HOST")
 	dbPort := os.Getenv("DB_PORT")
 	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&ssl-mode=REQUIRED", dbUser, dbPassword, dbHost, dbPort, dbName)
	
	port := os.Getenv("PORT")

	// Connect to database and migrate
	database.ConnectDb(connectionString)
	database.Migrate()

	// Initialize and run the router
	router := routes.InitRouter()
	portNum := fmt.Sprintf(":%s", port)
	router.Run(portNum)
}
