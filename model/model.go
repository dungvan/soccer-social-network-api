package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	*gorm.Model
	UserName string `gorm:"column:user_name;not null"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
	FullName string `gorm:"column:full_name;not null"`
	Birthday string `gorm:"column:birthday;not null"`
}

// HashAndSaltPassword encrypt password
func (u User) HashAndSaltPassword() User {
	hashAndSaltPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashAndSaltPassword)
	return u
}
