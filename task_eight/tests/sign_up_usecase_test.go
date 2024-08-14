package tests

import (
	"context"

	// "log"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/constants"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type signUpUseCaseSuite struct {
	suite.Suite
	repository *mocks.MockUserRepository
	usecase    domain.SignUpUseCase
	ctrl       *gomock.Controller
	ctx        context.Context
	ENV        *bootstrap.Env
}

func (suite *signUpUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockUserRepository(suite.ctrl)
	suite.usecase = usecases.NewSignUpUseCase(suite.repository)
	suite.ENV = bootstrap.NewEnv()
}

func (suite *signUpUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *signUpUseCaseSuite) createTestUser(errMess *domain.ErrorResponse) (domain.UserDTO, *domain.ErrorResponse) {
	userReq := domain.UserCreateRequest{
		Email:    constants.TestEmail,
		Password: constants.TestPassword,
	}

	suite.repository.
		EXPECT().
		CreateUser(gomock.Any(), userReq).
		Return(domain.UserDTO{}, errMess).
		Times(1)

	return suite.usecase.CreateUser(suite.ctx, userReq)
}

func (suite *signUpUseCaseSuite) TestCreateUser_Positive() {
	_, err := suite.createTestUser(nil)
	suite.Nil(err, "expected no error but got: %v", err)
}

func (suite *signUpUseCaseSuite) TestCreateUser_NilPointer_Negative() {
	suite.repository.EXPECT().
		CreateUser(gomock.Any(), domain.UserCreateRequest{}).
		Return(domain.UserDTO{}, &domain.ErrorResponse{Message: "Invalid request"}).
		Times(1)

	user, err := suite.usecase.CreateUser(suite.ctx, domain.UserCreateRequest{})

	suite.Error(err, "expected an error due to invalid input but got none")
	suite.Equal(domain.UserDTO{}, user, "expected an empty user object but got: %v", user)
	suite.Contains(err.Error(), "Invalid request", "expected specific error message")
}

func (suite *signUpUseCaseSuite) TestCreateUser_UserAlreadyExists() {
	// Create the first user
	_, firstErr := suite.createTestUser(nil)
	suite.Nil(firstErr, "expected no error when creating the first user but got: %v", firstErr)

	// Simulate user already exists error
	userAlreadyExistsError := &domain.ErrorResponse{Message: "User already exists"}
	_, err := suite.createTestUser(userAlreadyExistsError)
	suite.Error(err, "expected error since the user already exists but got none")
	suite.Contains(err.Error(), userAlreadyExistsError.Message, "expected specific error for user already exists but got: %v", err)
}

func (suite *signUpUseCaseSuite) TestCreateAccessToken() {
	user, err := suite.createTestUser(nil)
	suite.Nil(err, "Expected no error when creating the user but got")

	_, aErr := suite.usecase.CreateAccessToken(user, []byte(suite.ENV.JWT_SECRET))
	suite.Nil(aErr, "Expected no error while creating the access token")
}

func (suite *signUpUseCaseSuite) TestEncryptPassword() {
	password := "securepassword"

	hashedPassword, err := suite.usecase.EncryptPassword(password)
	suite.NoError(err, "Encryption should not return an error")

	suite.NotEmpty(hashedPassword, "Hashed password should not be empty")

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	suite.NoError(err, "Password comparison should succeed")

	anotherPassword := "anotherpassword"
	anotherHashedPassword, err := suite.usecase.EncryptPassword(anotherPassword)
	suite.NoError(err, "Encryption should not return an error")
	suite.NotEqual(hashedPassword, anotherHashedPassword, "Hashes for different passwords should not be the same")
}

func TestSignUp(t *testing.T) {
	suite.Run(t, new(signUpUseCaseSuite))
}
