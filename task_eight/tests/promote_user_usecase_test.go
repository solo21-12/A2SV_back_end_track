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

type promoteUserUseCaseSuite struct {
	suite.Suite
	repository *mocks.MockUserRepository
	usecase    domain.PromoteUseCase
	ctrl       *gomock.Controller
	ctx        context.Context
	ENV        *bootstrap.Env
}

func (suite *promoteUserUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockUserRepository(suite.ctrl)
	suite.usecase = usecases.NewPromoteUseCase(suite.repository)
	suite.ENV = bootstrap.NewEnv()
}

func (suite *promoteUserUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *promoteUserUseCaseSuite) TestPromoteUser_Success() {
	userID := "12345"

	suite.repository.
		EXPECT().
		PromoteUser(suite.ctx, userID).
		Return(nil) 

	err := suite.usecase.PromoteUser(userID, suite.ctx)

	suite.Nil(err, "Expected no error while promoting user")
}

func (suite *promoteUserUseCaseSuite) TestPromoteUser_Failure() {
	userID := "12345"
	expectedError := &domain.ErrorResponse{Message: "Promotion failed"}

	suite.repository.
		EXPECT().
		PromoteUser(suite.ctx, userID).
		Return(expectedError)

	err := suite.usecase.PromoteUser(userID, suite.ctx)

	suite.Equal(expectedError, err, "Expected the error response to match the promotion failure")
}

func TestPromoteUserUseCase(t *testing.T) {
	suite.Run(t, new(promoteUserUseCaseSuite))
}
