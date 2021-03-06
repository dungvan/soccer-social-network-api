package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	*gorm.Model
	UserName  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Birthday  *time.Time
	City      string
	Country   string
	About     string
	Quote     string
	Role      string `gorm:"default:'user'"`
	Score     uint
}

// HashAndSaltPassword encrypt password
func (u User) HashAndSaltPassword() User {
	hashAndSaltPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashAndSaltPassword)
	return u
}

// CompareHashAndPassword map the password with current hash password of user
func (u User) CompareHashAndPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

// GetCustomClaims get customs claims
func (u User) GetCustomClaims() map[string]interface{} {
	claims := make(map[string]interface{})
	userclaim := struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
	}{
		ID:       u.ID,
		Email:    u.Email,
		UserName: u.UserName,
	}
	claims["user"] = userclaim
	return claims
}

// GetIdentifier get identifier function
func (u User) GetIdentifier() uint {
	return u.ID
}
