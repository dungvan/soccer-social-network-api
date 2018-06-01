package tournament

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
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Team1       Team   `json:"team1"`
	Team2       Team   `json:"team2"`
	Team1Goals  *uint  `json:"team1_goals"`
	Team2Goals  *uint  `json:"team2_goals"`
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
