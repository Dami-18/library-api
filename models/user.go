package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

// Helper methods
func (user *User) HashPasswd(password string) error { // return type error
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPasswd(providedPasswd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPasswd))
	if err != nil {
		return err
	}
	return nil
}