package model

import (
	"github.com/jinzhu/gorm"
)

// TeamPlayer table struct
type TeamPlayer struct {
	*gorm.Model
	TeamID   uint
	UserID   uint
	Position string
}
