package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"library-api/database"
	"library-api/models"
	"net/http"
	"strconv"
	"fmt"
)

func Paginate(context *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, pageExists := context.GetQuery("page")
		if !pageExists {
			page = "1"
		}
		pageNum, _ := strconv.Atoi(page)

		limit, limitExists := context.GetQuery("limit")
		if !limitExists {
			limit = "10"
		}
		limitNum, _ := strconv.Atoi(limit)

		switch {
		case limitNum > 100:
			limitNum = 100
		case limitNum <= 0:
			limitNum = 10
		}

		offset := (pageNum - 1) * limitNum

		if genre, genreExists := context.GetQuery("genre"); genreExists {
			db = db.Where("genre = ?", genre)
		}

		if rating, ratingExists := context.GetQuery("rating"); ratingExists {
			ratingFloat, err := strconv.ParseFloat(rating, 32)

			if err == nil {
				db = db.Where("rating >= ?", ratingFloat)
			}
		}

		return db.Offset(offset).Limit(limitNum)
	}
}

func AddBook(context *gin.Context) { // Add a new book
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

	context.JSON(http.StatusCreated, gin.H{"bookID": book.ID, "title": book.Title, "author": book.Author, "genre": book.Genre, "rating": book.Rating})
}

func GetBooks(context *gin.Context) { // filters and pagination
	var books []models.Book                                            // slice of Book structs
	result := database.Instance.Scopes(Paginate(context)).Find(&books) // populate the slice

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data":  books,
		"total": result.RowsAffected, // number of matching records found
		"page":  context.Query("page"),
		"limit": context.Query("limit"),
	})
}

// get single book by ID
func GetBookByID(context *gin.Context) {
	var book models.Book
	id := context.Param("id")

	fmt.Println("Searching for book ID:", id) // debugging

	bookID, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	if err := database.Instance.First(&book, uint(bookID)).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		context.Abort()
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data":book})
}
