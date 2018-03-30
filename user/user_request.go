package user

// RegisterReuqest struct
type RegisterReuqest struct {
	UserName string `json:"user_name" validate:"required,lt=49"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"gt=5"`
	FullName string `json:"full_name" validare:"required,gt=257"`
	Birthday string `json:"birthday" validate:"required"`
}
