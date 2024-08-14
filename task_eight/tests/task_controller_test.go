package tests

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var fixedTime = time.Date(2024, time.March, 26, 10, 20, 20, 20, time.Local)

type TaskControllerSuite struct {
	suite.Suite
	usecase    *mocks.MockTaskUseCase
	controller *controllers.TaskController
	server     *httptest.Server
	ctrl       *gomock.Controller
}

func (suite *TaskControllerSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.usecase = mocks.NewMockTaskUseCase(suite.ctrl)

	suite.controller = &controllers.TaskController{
		TaskUseCase: suite.usecase,
	}

	router := gin.Default()
	router.POST("/tasks", suite.controller.Create)
	router.GET("/tasks", suite.controller.GetAll)
	router.GET("/tasks/:id", suite.controller.Get)
	router.DELETE("/tasks/:id", suite.controller.Delete)
	router.PUT("/tasks/:id", suite.controller.Update)
	suite.server = httptest.NewServer(router)
}

func (suite *TaskControllerSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *TaskControllerSuite) TestCreateTask_Success() {
	taskCreate := domain.TaskCreateDTO{
		Title:       "New Task",
		Description: "Description",
		DueDate:     fixedTime,
		Status:      "pending",
	}
	taskDTO := domain.TaskDTO{
		ID:          primitive.NewObjectID(),
		Title:       "New Task",
		Description: "Description",
		DueDate:     fixedTime,
		Status:      "pending",
	}

	suite.usecase.
		EXPECT().
		CreateTask(taskCreate, gomock.Any()).
		Return(taskDTO, nil).
		Times(1)

	inputJSON, err := json.Marshal(taskCreate)
	suite.NoError(err)

	response, err := http.Post(suite.server.URL+"/tasks", "application/json", bytes.NewBuffer(inputJSON))
	suite.NoError(err)

	defer response.Body.Close()

	// Check the response status
	suite.Equal(http.StatusCreated, response.StatusCode)
}

func (suite *TaskControllerSuite) TestCreateTask_Failure() {

	taskCreate := domain.TaskCreateDTO{
		Title:       "New Task",
		Description: "Description",
		DueDate:     fixedTime,
		Status:      "pending",
	}

	inputJSON, err := json.Marshal(taskCreate)
	suite.NoError(err)

	suite.usecase.
		EXPECT().
		CreateTask(gomock.Any(), gomock.Any()).
		Return(domain.TaskDTO{}, &domain.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad Request"})

	resp, err := http.Post(suite.server.URL+"/tasks", "application/json", bytes.NewBuffer(inputJSON))
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (suite *TaskControllerSuite) TestGetAllTasks_Success() {

	tasks := []domain.TaskDTO{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1", DueDate: fixedTime, Status: "pending"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2", DueDate: fixedTime, Status: "completed"},
	}

	suite.usecase.
		EXPECT().
		GetTasks(gomock.Any()).
		Return(tasks, nil)

	resp, err := http.Get(suite.server.URL + "/tasks")
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *TaskControllerSuite) TestGetAllTasks_Failure() {
	suite.usecase.
		EXPECT().
		GetTasks(gomock.Any()).
		Return(nil, &domain.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal Server Error"})

	resp, err := http.Get(suite.server.URL + "/tasks")
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(http.StatusInternalServerError, resp.StatusCode)
}

func (suite *TaskControllerSuite) TestGetTask_Failure() {

	objID := primitive.NewObjectID()
	suite.usecase.
		EXPECT().
		GetTaskByID(objID.Hex(), gomock.Any()).
		Return(domain.TaskDTO{}, &domain.ErrorResponse{Code: http.StatusNotFound, Message: "Not Found"})

	resp, err := http.Get(suite.server.URL + "/tasks/" + objID.Hex())
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}

func (suite *TaskControllerSuite) TestDeleteTask_Success() {

	objID := primitive.NewObjectID()
	suite.usecase.
		EXPECT().
		DeleteTask(objID.Hex(), gomock.Any()).
		Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, suite.server.URL+"/tasks/"+objID.Hex(), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(http.StatusNoContent, resp.StatusCode)
}

func (suite *TaskControllerSuite) TestDeleteTask_Failure() {
	objID := primitive.NewObjectID()

	suite.usecase.
		EXPECT().
		DeleteTask(objID.Hex(), gomock.Any()).
		Return(&domain.ErrorResponse{Code: http.StatusNotFound, Message: "Not Found"})

	req, _ := http.NewRequest(http.MethodDelete, suite.server.URL+"/tasks/"+objID.Hex(), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
