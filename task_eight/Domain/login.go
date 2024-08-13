package domain

import (
	"context"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        UserDTO `json:"user"`
	AccessToken string  `json:"accessToken"`
}

type LoginUseCase interface {
	GetUserEmail(ctx context.Context, email string) (User, *ErrorResponse)
	CreateAccessToken(user UserDTO, secret []byte) (accessToken string, err error)
	ValidatePassword(password string, userPassword string) bool
	GetJwtSecret() ([]byte, error)
}
