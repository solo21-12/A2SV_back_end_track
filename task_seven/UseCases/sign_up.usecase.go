package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

type signUpUseCase struct {
	userRespository domain.UserRepository
}

func NewSignUpUseCase(userRespository domain.UserRepository) domain.SignUpUseCase {
	return &signUpUseCase{
		userRespository: userRespository,
	}
}

func (s *signUpUseCase) CreateUser(ctx context.Context, user domain.UserCreateRequest) (domain.UserDTO, *domain.ErrorResponse) {
	return s.userRespository.CreateUser(ctx, user)
}
