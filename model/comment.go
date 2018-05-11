package model

import "github.com/jinzhu/gorm"

// Comment table struct
type Comment struct {
	*gorm.Model
	PostID    uint
	UserID    uint
	Content   string
	StarCount *StarCount `gorm:"polymorphic:Owner;"`
}
