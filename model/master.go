package model

import (
	"github.com/jinzhu/gorm"
)

// Master table struct
type Master struct {
	*gorm.Model
	UserID    uint
	OwnerID   uint
	OwnerType string
}
