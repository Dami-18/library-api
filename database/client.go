package database

import (
	"library-api/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB // Instance will be globally imported along with its methods, so don't change name `Instance`
var dbError error

func ConnectDb(connectionString string){
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if dbError != nil{
		log.Fatal(dbError)
		panic("Error connecting to database!")
	}
	log.Println("Connected to Database!")
}

func Migrate(){
	Instance.AutoMigrate(&models.User{},&models.Book{}) // both tables 
	log.Println("Database migration complete!")
}

