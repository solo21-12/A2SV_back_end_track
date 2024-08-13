package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
)

type loginUseCase struct {
	userRespository domain.UserRepository
}

func NewLoginUseCase(userRespository domain.UserRepository) domain.LoginUseCase {
	return &loginUseCase{
		userRespository: userRespository,
	}
}

func (l *loginUseCase) GetUserEmail(ctx context.Context, email string) (*domain.User, *domain.ErrorResponse) {
	return l.userRespository.GetUserEmail(ctx, email)

}
func (l *loginUseCase) CreateAccessToken(user domain.UserDTO, secret []byte) (accessToken string, err error) {
	return infrastructure.CreateAccessToken(user, secret)

}
func (l *loginUseCase) ValidatePassword(password string, userPassword string) bool {
	return infrastructure.ValidatePassword(password, userPassword)
}

func (l *loginUseCase) GetJwtSecret() ([]byte, error) {
	return infrastructure.GetJwtSecret()
}
