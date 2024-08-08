package model

type Auth struct {
	User  UserLogin `json:"user"`
	Token string    `json:"token"`
}
