package tournament

import "time"

// IndexRequest struct
type IndexRequest struct {
	page uint `form:"page"`
}

// CreateRequest struct
type CreateRequest struct {
	UserID      uint      `validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Teams       []ReqTeam `json:"teams" validate:"dive"`
}

// ReqTeam struct
type ReqTeam struct {
	ID    uint   `json:"id" validate:"required"`
	Group string `json:"group" validate:"omitempty,required"`
}
