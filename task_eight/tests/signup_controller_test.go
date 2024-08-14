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
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignUpControllerTestSuite struct {
	suite.Suite
	usecase    *mocks.MockSignUpUseCase
	controller controllers.SignupController
	server     *httptest.Server
	ctrl       *gomock.Controller
	ENV        bootstrap.Env
}

func (suite *SignUpControllerTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.usecase = mocks.NewMockSignUpUseCase(suite.ctrl)
	suite.ENV = *bootstrap.NewEnv()

	suite.controller = controllers.SignupController{
		SignupUsecase: suite.usecase,
	}

	router := gin.Default()
	router.POST("/register", suite.controller.SignUp)
	testingServer := httptest.NewServer(router)
	suite.server = testingServer
}

func (suite *SignUpControllerTestSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *SignUpControllerTestSuite) TestSignUp_Success() {
	// Define input and expected output

	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	input := domain.UserCreateRequest{
		Email:    "testuser",
		Password: string(encryptPassword),
	}

	expectedUserDTO := domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: "testuser",
		Role:  "admin",
	}

	accessToken, ne := infrastructure.CreateAccessToken(expectedUserDTO, []byte(suite.ENV.JWT_SECRET))

	suite.Nil(ne, "expecte no error while creating the access token")

	suite.usecase.
		EXPECT().
		CreateAccessToken(expectedUserDTO, gomock.Eq([]byte(suite.ENV.JWT_SECRET))).
		Return(accessToken, nil)

	suite.usecase.
		EXPECT().
		EncryptPassword(input.Password).
		Return(string(encryptPassword), nil).
		Times(1)

	suite.usecase.
		EXPECT().
		GetJwtSecret().
		Return([]byte(suite.ENV.JWT_SECRET), nil).
		Times(1)

	suite.usecase.
		EXPECT().
		CreateUser(gomock.Any(), input).
		Return(expectedUserDTO, nil).
		Times(1)

	// Convert input to JSON
	inputJSON, err := json.Marshal(input)
	suite.NoError(err)

	// Send HTTP POST request to the server
	response, err := http.Post(suite.server.URL+"/register", "application/json", bytes.NewBuffer(inputJSON))
	suite.NoError(err)

	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusCreated, response.StatusCode)
}

func TestSignUpControllerTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpControllerTestSuite))
}
