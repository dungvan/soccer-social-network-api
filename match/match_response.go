package match

import "time"

// IndexResponse struct
type IndexResponse struct {
	Total   uint        `json:"total"`
	Matches []RespMatch `json:"matches"`
}

// CreateResponse struct
type CreateResponse struct {
	MatchID uint `json:"match_id"`
}

// RespMatch struct
type RespMatch struct {
	TypeOfStatusCode int             `json:"-"`
	ID               uint            `json:"id"`
	Tournament       *RespTournament `json:"tournament"`
	Description      string          `json:"description"`
	Master           RespMaster      `json:"master"`
	StartDate        time.Time       `json:"start_date"`
	Team1            RespTeam        `json:"team1"`
	Team2            RespTeam        `json:"team2"`
	Team1Goals       *uint           `json:"team1_goals"`
	Team2Goals       *uint           `json:"team2_goals"`
}

// RespMaster struct
type RespMaster struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// RespTeam struct
type RespTeam struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Master      RespMaster `json:"master"`
	Players     []Player   `json:"players"`
}

// RespTournament struct
type RespTournament struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
