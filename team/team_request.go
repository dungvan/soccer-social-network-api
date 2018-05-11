package team

// IndexRequest struct
type IndexRequest struct {
	Page uint `form:"page" validate:"omitempty,min=1"`
}

// CreateRequest struct
type CreateRequest struct {
	UserID      uint            `validate:"required"`
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description" validate:"required"`
	Players     []PlayerRequest `json:"players" validate:"required,max=16,dive"`
}

// UpdateRequest struct
type UpdateRequest struct {
	ID          uint            `json:"id" validate:"required,min=1"`
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description" validate:"required"`
	PlayersAdd  []PlayerRequest `json:"players_add" validate:"required,dive"`
	PlayersDel  []uint          `json:"players_del" validate:"required"`
}

// PlayerRequest struct
type PlayerRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Position string `json:"position" validate:"required,eq=gk|eq=sw|eq=cb|eq=lb|eq=rb|eq=dm|eq=lwb|eq=rwb|eq=cm|eq=am|eq=lw|eq=rw|eq=wf|eq=cf|eq=any"`
}
