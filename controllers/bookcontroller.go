package controllers

import (
	"library-api/database"
	"library-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddBook(context *gin.Context){ // Add a new book
	var book models.Book
	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if book.Title == "" || book.Author == "" || book.Genre == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Title, author and genre are required!"})
		context.Abort()
        return
    }

	if book.Rating != 0 && (book.Rating < 1 || book.Rating > 5) {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		context.Abort()
        return
    }

	record := database.Instance.Create(&book)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error in adding the book!"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"bookID":book.ID, "title":book.Title, "author":book.Author, "genre":book.Author, "rating":book.Rating})
}

