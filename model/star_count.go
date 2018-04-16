package model

import "github.com/jinzhu/gorm"

// StarCount table struct
type StarCount struct {
	*gorm.Model
	Quantity  uint
	OwnerID   uint
	OwnerType string
}
