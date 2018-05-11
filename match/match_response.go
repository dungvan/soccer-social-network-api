package match

import "time"

// CreateResponse struct
type CreateResponse struct {
	MatchID uint `json:"match_id"`
}

// RespMatch struct
type RespMatch struct {
	TypeOfStatusCode int        `json:"-"`
	ID               uint       `json:"id"`
	TournamentID     *uint      `json:"tournament_id"`
	Description      string     `json:"description"`
	Master           RespMaster `json:"master"`
	StartDate        time.Time  `json:"start_date"`
	Team1ID          uint       `json:"team1_id"`
	Team2ID          uint       `json:"team2_id"`
}

// RespMaster struct
type RespMaster struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
