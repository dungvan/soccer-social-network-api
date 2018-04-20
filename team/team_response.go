package team

// CreateResponse struct
type CreateResponse struct {
	TeamID uint `json:"team_id"`
}

// IndexResponse struct
type IndexResponse struct {
	Master RespTeams `json:"master"`
	Player RespTeams `json:"player"`
}

// RespTeams struct
type RespTeams struct {
	ResultCount int        `json:"result_count"`
	Teams       []RespTeam `json:"teams"`
}

// RespTeam struct
type RespTeam struct {
	TypeOfStatusCode int        `json:"-"`
	ID               uint       `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Master           RespMaster `json:"master"`
	Players          []Player   `json:"players"`
}

// RespMaster struct
type RespMaster struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
}
