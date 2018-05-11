package team

// CreateResponse struct
type CreateResponse struct {
	TeamID uint `json:"team_id"`
}

// ByUserResponse struct
type ByUserResponse struct {
	Master IndexResponse `json:"master"`
	Player IndexResponse `json:"player"`
}

// IndexResponse struct
type IndexResponse struct {
	TypeOfStatusCode int        `json:"-"`
	Total            uint       `json:"total"`
	Teams            []RespTeam `json:"teams"`
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
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
