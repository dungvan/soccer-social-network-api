package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Match table struct
type Match struct {
	*gorm.Model
	Master      *Master `gorm:"polymorphic:Owner"`
	Description string
	StartDate   time.Time
	Team1ID     uint
	Team2ID     uint
	Location    *Location
}
