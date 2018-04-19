package model

import (
	"github.com/jinzhu/gorm"
)

// Team table struct
type Team struct {
	*gorm.Model
	Name        string
	Description string
	Master      Master `gorm:"polymorphic:Owner"`
	Players     []User `gorm:"many2many:team_players"`
	MaxMembers  uint
}
