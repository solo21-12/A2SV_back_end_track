package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type loginControllerTestSuite struct {
	suite.Suite
	usecase    *mocks.MockLoginUseCase
	controller controllers.LoginController
	server     *httptest.Server
	ctrl       *gomock.Controller
	ENV        bootstrap.Env
}

func (suite *loginControllerTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.usecase = mocks.NewMockLoginUseCase(suite.ctrl)
	suite.ENV = *bootstrap.NewEnv()

	suite.controller = controllers.LoginController{
		LoginUseCase: suite.usecase,
		Env:          &suite.ENV,
	}

	router := gin.Default()
	router.POST("/login", suite.controller.Login)
	suite.server = httptest.NewServer(router)
}

func (suite *loginControllerTestSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *loginControllerTestSuite) sendLoginRequest(input domain.LoginRequest) (*http.Response, error) {
	inputJSON, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return http.Post(suite.server.URL+"/login", "application/json", bytes.NewBuffer(inputJSON))
}

func (suite *loginControllerTestSuite) setupMockGetJwtSecret() {
	suite.usecase.EXPECT().
		GetJwtSecret().
		Return([]byte(suite.ENV.JWT_SECRET), nil).
		Times(1)
}

func (suite *loginControllerTestSuite) TestLogin_Success() {
	// Setup
	plainPassword := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	input := domain.LoginRequest{
		Email:    "testuser@example.com",
		Password: plainPassword,
	}

	existingUser := domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: "testuser@example.com",
		Role:  "admin",
	}

	userWithPassword := domain.User{
		ID:       existingUser.ID,
		Email:    existingUser.Email,
		Role:     existingUser.Role,
		Password: string(hashedPassword),
	}

	accessToken, err := infrastructure.CreateAccessToken(existingUser, []byte(suite.ENV.JWT_SECRET))
	suite.NoError(err)

	// Mock expectations
	suite.setupMockGetJwtSecret()

	suite.usecase.EXPECT().
		GetUserEmail(gomock.Any(), input.Email).
		Return(&userWithPassword, nil).
		Times(1)

	suite.usecase.EXPECT().
		ValidatePassword(input.Password, string(hashedPassword)).
		Return(true).
		Times(1)

	suite.usecase.EXPECT().
		CreateAccessToken(existingUser, []byte(suite.ENV.JWT_SECRET)).
		Return(accessToken, nil).
		Times(1)

	// Send HTTP POST request to the server
	response, err := suite.sendLoginRequest(input)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusOK, response.StatusCode)

	// Decode the response body
	var responseLogin domain.LoginResponse
	err = json.NewDecoder(response.Body).Decode(&responseLogin)
	suite.NoError(err)
	suite.Equal(existingUser, responseLogin.User)
	suite.Equal(accessToken, responseLogin.AccessToken)
}

func (suite *loginControllerTestSuite) TestLogin_InvalidDataFormat() {
	// Setup
	suite.setupMockGetJwtSecret()

	// Send HTTP POST request with invalid data
	response, err := http.Post(suite.server.URL+"/login", "application/json", bytes.NewBuffer([]byte("invalid data")))
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *loginControllerTestSuite) TestLogin_UserNotFound() {
	// Setup
	input := domain.LoginRequest{
		Email:    "nonexistentuser@example.com",
		Password: "password123",
	}

	suite.setupMockGetJwtSecret()

	suite.usecase.EXPECT().
		GetUserEmail(gomock.Any(), input.Email).
		Return(&domain.User{}, &domain.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}).
		Times(1)

	// Send HTTP POST request to the server
	response, err := suite.sendLoginRequest(input)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusNotFound, response.StatusCode)
}

func (suite *loginControllerTestSuite) TestLogin_InvalidPassword() {
	// Setup
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("wrongpassword"), bcrypt.DefaultCost)
	input := domain.LoginRequest{
		Email:    "testuser@example.com",
		Password: "wrongpassword",
	}

	existingUser := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "testuser@example.com",
		Role:     "admin",
		Password: string(hashedPassword),
	}

	// Mock expectations
	suite.setupMockGetJwtSecret()

	suite.usecase.EXPECT().
		GetUserEmail(gomock.Any(), input.Email).
		Return(&existingUser, nil).
		Times(1)

	suite.usecase.EXPECT().
		ValidatePassword(input.Password, existingUser.Password).
		Return(false).
		Times(1)

	// Send HTTP POST request to the server
	response, err := suite.sendLoginRequest(input)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *loginControllerTestSuite) TestLogin_TokenCreationError() {
	// Setup
	plainPassword := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	input := domain.LoginRequest{
		Email:    "testuser@example.com",
		Password: plainPassword,
	}

	existingUser := domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: "testuser@example.com",
		Role:  "admin",
	}
	userWithPassword := domain.User{
		ID:       existingUser.ID,
		Email:    existingUser.Email,
		Role:     existingUser.Role,
		Password: string(hashedPassword),
	}

	// Mock expectations
	suite.usecase.EXPECT().
		GetUserEmail(gomock.Any(), input.Email).
		Return(&userWithPassword, nil).
		Times(1)

	suite.setupMockGetJwtSecret()

	suite.usecase.EXPECT().
		ValidatePassword(input.Password, string(hashedPassword)).
		Return(true).
		Times(1)

	suite.usecase.EXPECT().
		CreateAccessToken(existingUser, []byte(suite.ENV.JWT_SECRET)).
		Return("", &domain.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Token creation error",
		}).
		Times(1)

	// Send HTTP POST request to the server
	response, err := suite.sendLoginRequest(input)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

func TestLoginControllerTestSuite(t *testing.T) {
	suite.Run(t, new(loginControllerTestSuite))
}
