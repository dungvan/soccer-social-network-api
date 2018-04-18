package model

import (
	"github.com/jinzhu/gorm"
)

// Image table struct
type Image struct {
	*gorm.Model
	PostID uint
	Name   string
}
