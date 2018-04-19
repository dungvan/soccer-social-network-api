package model

import (
	"github.com/jinzhu/gorm"
)

// Tournament table struct
type Tournament struct {
	*gorm.Model
	Name        string
	Description string
	Master      []Master `gorm:"polymorphic:Owner"`
}
