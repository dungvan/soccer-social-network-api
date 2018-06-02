package tournament

import (
	"time"
)

// Team struct
type Team struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Master      RespMaster `json:"master"`
	Players     []Player   `json:"players"`
}

// Match struct
type Match struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	Team1ID     uint      `json:"team1_id"`
	Team2ID     uint      `json:"team2_id"`
	Team1Goals  *uint     `json:"team1_goals"`
	Team2Goals  *uint     `json:"team2_goals"`
}

// Player truct
type Player struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     uint   `json:"score"`
	Position  string `json:"position"`
}
