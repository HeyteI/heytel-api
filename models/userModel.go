package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is a struct model containing user data fields and methods for hashing and checking passwords 

type User struct {
	Base      `mapstructure:",squash"`
	FirstName string     `json:"firstname" binding:"required"`
	LastName  string     `json:"lastname" binding:"required"`
	Username  string     `json:"username" gorm:"unique"`
	Password  string     `json:"password"`
	Group     string     `json:"group" binding:"required"`
	Birthday  *time.Time `json:"birthday"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Photo_url string     `json:"photo" default:"https://ui-avatars.com/api/?name=H"`
}

// HashPassword hashes the password of a user
// and returns an error if any
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks if the password of a user is correct
// and returns an error if any
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
