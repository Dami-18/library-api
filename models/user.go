package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role string `json:"role"` // role based access, admin or user
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