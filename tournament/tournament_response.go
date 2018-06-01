package tournament

import "time"

// CreateResponse struct
type CreateResponse struct {
	TournamentID uint `json:"tournament_id"`
}

// IndexResponse struct
type IndexResponse struct {
	Total       uint             `json:"total"`
	Tournaments []RespTournament `json:"tournaments"`
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
	Matches          []Match    `json:"matches"`
}

// RespMaster struct
type RespMaster struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
