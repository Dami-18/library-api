package models

import (
	"gorm.io/gorm"
)

// book model
type Book struct {
	gorm.Model
	Title  string  `json:"title" gorm:"not null"`
	Author string  `json:"author" gorm:"not null"`
	Genre  string  `json:"genre" gorm:"not null"`
	Rating float32 `json:"rating" gorm:"default:0"`
}
