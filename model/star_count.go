package model

import (
	"github.com/jinzhu/gorm"
)

// StarCount table struct
type StarCount struct {
	*gorm.Model
	OwnerID   int
	OwnerType string
	Quantity  uint
}
