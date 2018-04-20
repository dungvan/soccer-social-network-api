package match

import (
	"time"
)

// CreateRequest struct
type CreateRequest struct {
	UserID      uint      `validate:"required"`
	Description string    `json:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	Team1ID     uint      `json:"team1_id" validate:"required"`
	Team2ID     uint      `json:"team2_id" validate:"required"`
}
