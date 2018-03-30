package model

import "github.com/jinzhu/gorm"

// Hashtag struct
type Hashtag struct {
	*gorm.Model
	Key string `gorm:"column:key;not null"`
}
