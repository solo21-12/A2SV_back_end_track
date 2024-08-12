package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PromoteResponse struct {
	Message string `json:"message"`
}

type PromoteUseCase interface {
	PromoteUser(userID primitive.ObjectID, ctx context.Context) *ErrorResponse
}
