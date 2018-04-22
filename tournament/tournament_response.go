package tournament

import "time"

// CreateResponse struct
type CreateResponse struct {
	TournamentID uint `json:"tournament_id"`
}

// RespTournament struct
type RespTournament struct {
	TypeOfStatusCode int        `json:"-"`
	ID               uint       `json:"id"`
	Description      string     `json:"description"`
	Master           RespMaster `json:"master"`
	StartDate        time.Time  `json:"start_date"`
	EndDate          time.Time  `json:"end_date"`
	Teams            []Team     `json:"teams"`
}

// RespMaster struct
type RespMaster struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
}
