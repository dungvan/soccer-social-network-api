package model

import (
	"github.com/jinzhu/gorm"
)

// LocationID type
type LocationID uint

// Post struct
type Post struct {
	*gorm.Model
	UserID     uint
	Caption    string
	Location   Location
	LocationID *LocationID
	Hashtags   []Hashtag `gorm:"many2many:post_hashtags"`
	Stars      []Star    `gorm:"many2many:post_stars"`
	StarCount  StarCount `gorm:"polymorphic:Owner"`
}
