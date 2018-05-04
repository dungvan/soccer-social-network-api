package user

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
