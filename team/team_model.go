package team

// Player truct
type Player struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Score    uint   `json:"score"`
	Position string `json:"position"`
}
