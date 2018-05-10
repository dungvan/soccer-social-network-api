package user

import "time"

// RegisterReuqest struct
type RegisterReuqest struct {
	UserName             string `json:"user_name" validate:"required,lt=49"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"gt=5"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	FirstName            string `json:"first_name" validate:"required,lt=257"`
	LastName             string `json:"last_name" validate:"required,lt=257"`
}

// LoginRequest struct
type LoginRequest struct {
	UserNameOrEmail string `json:"user_name_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

// FriendRequest struct
type FriendRequest struct {
	UserID     uint `json:"user_id" validate:"required"`
	UserFollow uint `json:"user_follow" validate:"required"`
}

// IndexRequest struct
type IndexRequest struct {
	Page uint `form:"page" validate:"omitempty,min=1"`
}

// UpdateRequest struct
type UpdateRequest struct {
	updateField
	Password             string `json:"password" validate:"omitempty,gt=5"`
	PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password"`
}

type updateField struct {
	ID        uint       `json:"id" validate:"required"`
	FirstName string     `json:"first_name" validate:"gt=0"`
	LastName  string     `json:"last_name" validate:"gt=0"`
	City      string     `json:"city"`
	Country   string     `json:"country"`
	About     string     `json:"About"`
	Quote     string     `json:"quote"`
	Birthday  *time.Time `json:"birthday"`
}
