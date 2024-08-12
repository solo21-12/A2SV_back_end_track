package domain

import (
	"context"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	User        UserDTO `json:"user"`
	AccessToken string  `json:"accessToken"`
}

type SignUpUseCase interface {
	CreateUser(ctx context.Context, user UserCreateRequest) (UserDTO, *ErrorResponse)
	CreateAccessToken(user UserDTO, secret []byte) (accessToken string, err error)
	EncryptPassword(password string) (string, error)
	GetJwtSecret() ([]byte, error)
}
