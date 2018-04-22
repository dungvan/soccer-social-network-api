package match

import (
	"time"
)

// CreateRequest struct
type CreateRequest struct {
	TournamentID uint      `json:"tourmanent_id" validate:"omitempty,required"`
	UserID       uint      `validate:"required"`
	Description  string    `json:"description" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"required"`
	Team1ID      uint      `json:"team1_id" validate:"required,nefield=Team2ID"`
	Team2ID      uint      `json:"team2_id" validate:"required,nefield=Team1ID"`
}
