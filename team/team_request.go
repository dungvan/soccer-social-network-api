package team

// IndexRequest struct
type IndexRequest struct {
}

// CreateRequest struct
type CreateRequest struct {
	MasterID    uint            `json:"master_id" validate:"required"`
	Description string          `json:"description"`
	Players     []PlayerRequest `json:"players" validate:"required,max=16,dive"`
}

// PlayerRequest struct
type PlayerRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Position string `json:"position" validate:"required,eq="`
}
