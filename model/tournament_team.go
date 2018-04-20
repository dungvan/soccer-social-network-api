package model

import (
	"github.com/jinzhu/gorm"
)

// TournamentTeam table struct
type TournamentTeam struct {
	*gorm.Model
	TournamentID uint
	TeamID       uint
	Group        string
	Score        uint
}
