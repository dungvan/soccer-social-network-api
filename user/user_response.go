package user

import "time"

// LoginResponse response.
type LoginResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

// IndexResponse struct
type IndexResponse struct {
	TypeOfStatusCode int        `json:"-"`
	Total            uint       `json:"total"`
	Users            []RespUser `json:"users"`
}

// RespUser struct
type RespUser struct {
	TypeOfStatusCode int        `json:"-"`
	ID               uint       `json:"id"`
	Email            string     `json:"email"`
	UserName         string     `json:"user_name"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	Birthdate        *time.Time `json:"birthday"`
	About            string     `json:"about"`
	Quote            string     `json:"quote"`
	City             string     `json:"city"`
	Country          string     `json:"country"`
	Role             string     `json:"role"`
}
