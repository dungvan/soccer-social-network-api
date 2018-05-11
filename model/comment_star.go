package model

import (
	"github.com/jinzhu/gorm"
)

// CommentStar table struct.
type CommentStar struct {
	*gorm.Model
	CommentID uint
	UserID    uint
}
