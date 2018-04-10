package model

import (
	"github.com/jinzhu/gorm"
)

// Location struct
type Location struct {
	*gorm.Model
	PlaceID   string
	PostCount uint
	Post      []Post
}
