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

// PlayerRequest struct
type PlayerRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Position string `json:"position" validate:"required,eq=gk|eq=sw|eq=cb|eq=lb|eq=rb|eq=dm|eq=lwb|eq=rwb|eq=cm|eq=am|eq=lw|eq=rw|eq=wf|eq=cf|eq=any"`
}
