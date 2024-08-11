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
	
}
