package model

import (
	"github.com/jinzhu/gorm"
)

// PostStar table struct.
type PostStar struct {
	*gorm.Model
	PostID uint
	UserID uint
}
