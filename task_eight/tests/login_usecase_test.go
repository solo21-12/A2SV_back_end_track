package tests

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
)

const (
	testEmail       = "test@gmail.com"
	testPassword    = "strongpassword"
	invalidEmail    = "foo@gmail.com"
	invalidPassword = "wrongpassword"
)

type loginUseCaseSuite struct {
	suite.Suite
	repository    *mocks.MockUserRepository
	signupUseCase domain.SignUpUseCase
	usecase       domain.LoginUseCase
	ctrl          *gomock.Controller
	ctx           context.Context
	ENV           *bootstrap.Env
}

func (suite *loginUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockUserRepository(suite.ctrl)
	suite.signupUseCase = usecases.NewSignUpUseCase(suite.repository)
	suite.usecase = usecases.NewLoginUseCase(suite.repository)
	suite.ENV = bootstrap.NewEnv()
}

func (suite *loginUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *loginUseCaseSuite) createTestUser(errMess *domain.ErrorResponse) (domain.UserDTO, *domain.ErrorResponse) {
	userReq := domain.UserCreateRequest{
		Email:    testEmail,
		Password: testPassword,
	}

	suite.repository.EXPECT().
		CreateUser(gomock.Any(), userReq).
		Return(domain.UserDTO{}, errMess).
		Times(1)

	return suite.signupUseCase.CreateUser(suite.ctx, userReq)
}

func (suite *loginUseCaseSuite) getUser() (*domain.User, *domain.ErrorResponse) {
	suite.repository.
		EXPECT().
		GetUserEmail(gomock.Any(), testEmail).
		Return(&domain.User{Email: testEmail}, nil)

	return suite.usecase.GetUserEmail(suite.ctx, testEmail)
}

func (suite *loginUseCaseSuite) validatePassword(password, hashedPassword string, expected bool) {
	valid := suite.usecase.ValidatePassword(password, hashedPassword)
	suite.Equal(expected, valid, "Password validation result mismatch")
}

func (suite *loginUseCaseSuite) hashPassword(password string) (string, error) {
	return suite.signupUseCase.EncryptPassword(password)
}

func (suite *loginUseCaseSuite) TestGetUserEmail() {
	_, err := suite.createTestUser(nil)
	suite.Nil(err, "Error creating user")

	_, retErr := suite.getUser()
	suite.Nil(retErr, "Error retrieving user")

}

func (suite *loginUseCaseSuite) TestGetUserEmail_NotFound() {
	suite.repository.
		EXPECT().
		GetUserEmail(gomock.Any(), invalidEmail).
		Return(&domain.User{}, &domain.ErrorResponse{Message: "User not found"})

	retrievedUser, retErr := suite.usecase.GetUserEmail(suite.ctx, invalidEmail)

	suite.Error(retErr, "Expected error not received")
	suite.Equal(&domain.User{}, retrievedUser, "Retrieved user should be empty")
	suite.Contains(retErr.Message, "User not found", "Error message mismatch")
}

func (suite *loginUseCaseSuite) TestCreateAccessToken() {
	user, err := suite.createTestUser(nil)
	suite.Nil(err, "Error creating user")

	_, aErr := suite.usecase.CreateAccessToken(user, []byte(suite.ENV.JWT_SECRET))
	suite.Nil(aErr, "Error creating access token")
}

func (suite *loginUseCaseSuite) TestValidatePassword_Valid() {
	hashedPassword, err := suite.hashPassword(testPassword)
	suite.Nil(err, "Error hashing password")

	suite.validatePassword(testPassword, hashedPassword, true)
}

func (suite *loginUseCaseSuite) TestValidatePassword_Invalid() {
	hashedPassword, err := suite.hashPassword(testPassword)
	suite.Nil(err, "Error hashing password")

	suite.validatePassword(invalidPassword, hashedPassword, false)
}

func TestLoginUseCase(t *testing.T) {
	suite.Run(t, new(loginUseCaseSuite))
}
