package model

import "github.com/jinzhu/gorm"

// Comment table struct
type Comment struct {
	*gorm.Model
	PostID    uint
	StarCount *StarCount `gorm:"polymorphic:Owner;"`
}
