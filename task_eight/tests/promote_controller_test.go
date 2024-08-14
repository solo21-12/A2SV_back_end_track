package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
)

type promoteControllerTestSuite struct {
	suite.Suite
	usecase    *mocks.MockPromoteUseCase
	controller controllers.PromoteController
	server     *httptest.Server
	ctrl       *gomock.Controller
}

func (suite *promoteControllerTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.usecase = mocks.NewMockPromoteUseCase(suite.ctrl)
	suite.controller = controllers.PromoteController{
		PromoteUseCase: suite.usecase,
		Env:            &bootstrap.Env{}, // Use an actual Env setup if needed
	}

	router := gin.Default()
	router.POST("/promote/:id", suite.controller.PromoteUser)
	suite.server = httptest.NewServer(router)
}

func (suite *promoteControllerTestSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *promoteControllerTestSuite) TestPromoteUser_Success() {
	userID := "user123"

	suite.usecase.EXPECT().
		PromoteUser(userID, gomock.Any()).
		Return(nil).
		Times(1)

	// Send HTTP POST request to the server
	response, err := http.Post(suite.server.URL+"/promote/"+userID, "application/json", nil)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusOK, response.StatusCode)

	// Decode the response body
	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err)
	suite.Equal("User promoted successfully", responseBody["message"])
}

func (suite *promoteControllerTestSuite) TestPromoteUser_UserNotFound() {
	userID := "nonexistentuser"

	suite.usecase.EXPECT().
		PromoteUser(userID, gomock.Any()).
		Return(&domain.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}).
		Times(1)

	// Send HTTP POST request to the server
	response, err := http.Post(suite.server.URL+"/promote/"+userID, "application/json", nil)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusNotFound, response.StatusCode)

	// Decode the response body
	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err)
	suite.Equal("User not found", responseBody["error"])
}

func (suite *promoteControllerTestSuite) TestPromoteUser_Error() {
	userID := "user123"

	suite.usecase.EXPECT().
		PromoteUser(userID, gomock.Any()).
		Return(&domain.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Promotion error",
		}).
		Times(1)

	// Send HTTP POST request to the server
	response, err := http.Post(suite.server.URL+"/promote/"+userID, "application/json", nil)
	suite.NoError(err)
	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusInternalServerError, response.StatusCode)

	// Decode the response body
	var responseBody map[string]string
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err)
	suite.Equal("Promotion error", responseBody["error"])
}

func TestPromoteControllerTestSuite(t *testing.T) {
	suite.Run(t, new(promoteControllerTestSuite))
}
