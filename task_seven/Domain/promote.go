package domain

import (
	"context"

)

type PromoteResponse struct {
	Message string `json:"message"`
}

type PromoteUseCase interface {
	PromoteUser(userID string, ctx context.Context) *ErrorResponse
}
