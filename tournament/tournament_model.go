package tournament

// Team struct
type Team struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
	Score uint   `json:"score"`
}
