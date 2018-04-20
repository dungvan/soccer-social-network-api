package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tournament table struct
type Tournament struct {
	*gorm.Model
	Name        string
	Description string
	Master      *Master `gorm:"polymorphic:Owner"`
	Teams       []Team  `gorm:"many2many:tourmanent_teams"`
	StartDate   time.Time
	EndDate     time.Time
}
