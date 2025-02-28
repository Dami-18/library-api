package main

import (
	"OpenSoft-MT/database"
	"OpenSoft-MT/routes"
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

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	database.ConnectDb(connectionString) 
    database.Migrate()

	router := routes.InitRouter()
	port := os.Getenv("PORT")
	portNum := fmt.Sprintf(":%s",port)
	router.Run(portNum)
}