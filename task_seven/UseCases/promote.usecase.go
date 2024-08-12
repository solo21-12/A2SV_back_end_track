package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type promoteUseCase struct {
	userRepository domain.UserRepository
}

func NewPromoteUseCase(useRepository domain.UserRepository) domain.PromoteUseCase {
	return &promoteUseCase{
		userRepository: useRepository,
	}
}

func (u *promoteUseCase) PromoteUser(userID primitive.ObjectID, ctx context.Context) *domain.ErrorResponse {
	return u.userRepository.PromoteUser(ctx, userID)
}
