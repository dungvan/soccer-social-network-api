package user

// LoginResponse response.
type LoginResponse struct {
	Token string `json:"token"`
}

// RespUser struct
type RespUser struct {
	TypeOfStatusCode int    `json:"-"`
	ID               uint   `json:"id"`
	Email            string `json:"email"`
	UserName         string `json:"user_name"`
	Fullname         string `json:"full_name"`
	Birthday         string `json:"birthday"`
}
