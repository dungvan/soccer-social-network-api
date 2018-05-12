package team

// Player truct
type Player struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     uint   `json:"score"`
	Position  string `json:"position"`
}
