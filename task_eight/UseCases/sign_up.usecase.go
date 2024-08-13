package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
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

func (s *signUpUseCase) CreateAccessToken(user domain.UserDTO, secret []byte) (accessToken string, err error) {
	return infrastructure.CreateAccessToken(user, secret)
}

func (s *signUpUseCase) EncryptPassword(password string) (string, error) {
	return infrastructure.EncryptPassword(password)
}


func (s *signUpUseCase) GetJwtSecret() ([]byte, error){
	return infrastructure.GetJwtSecret()
}