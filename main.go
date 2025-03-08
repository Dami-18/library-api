package main

import (
	"library-api/database"
	"library-api/routes"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
)
func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	dbString := os.Getenv("DATABASE_URL")

	connectionString := fmt.Sprintf(dbString)

	database.ConnectDb(connectionString) 
    database.Migrate()

	router := routes.InitRouter()
	port := os.Getenv("PORT")
	portNum := fmt.Sprintf(":%s",port)
	router.Run(portNum)
}