package model

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	*gorm.Model
	UserName string `gorm:"column:user_name;not null"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}
