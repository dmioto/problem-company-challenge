package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);" json:"firstName"`
	LastName  string `gorm:"type:varchar(255);" json:"lastName"`
	Email     string `gorm:"type:varchar(255);" json:"email"`
	Password  string `gorm:"type:varchar(255);" json:"password"`
}

// Return a instance of User
func GetUser() User {
	var user User
	return user
}

// Return a list of instance of user
func GetUsers() []User {
	var users []User
	return users
}
