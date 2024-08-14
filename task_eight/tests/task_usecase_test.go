package tests

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUseCaseSuite struct {
	suite.Suite
	repository *mocks.MockTaskRepository
	usecase    domain.TaskUseCase
	ctrl       *gomock.Controller
	ctx        context.Context
	ENV        *bootstrap.Env
}

func (suite *taskUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockTaskRepository(suite.ctrl)
	suite.usecase = usecases.NewTaskUseCase(suite.repository)
	suite.ENV = bootstrap.NewEnv()
}

func (suite *taskUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *taskUseCaseSuite) createTask() domain.TaskCreateDTO {
	newTask := domain.TaskCreateDTO{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "Pending",
	}

	return newTask
}

func (suite *taskUseCaseSuite) TestCreateTask_Success() {
	newTask := suite.createTask()

	createdTask := domain.TaskDTO{
		ID:          primitive.NewObjectID(),
		Title:       newTask.Title,
		Description: newTask.Description,
		DueDate:     newTask.DueDate,
		Status:      newTask.Status,
	}

	suite.repository.EXPECT().
		CreateTask(newTask, suite.ctx).
		Return(createdTask, nil).
		Times(1)

	result, err := suite.usecase.CreateTask(newTask, suite.ctx)
	suite.Nil(err, "Expected no error while creating the task")
	suite.Equal(createdTask, result, "Expected created task to match")
}

func (suite *taskUseCaseSuite) TestCreateTask_Failure() {
	newTask := suite.createTask()

	errResp := &domain.ErrorResponse{Message: "Creation failed"}
	
	suite.repository.EXPECT().
		CreateTask(newTask, suite.ctx).
		Return(domain.TaskDTO{}, errResp).
		Times(1)

	result, err := suite.usecase.CreateTask(newTask, suite.ctx)
	suite.Equal(errResp, err, "Expected error response while creating the task")
	suite.Equal(domain.TaskDTO{}, result, "Expected empty task DTO on error")
}

func (suite *taskUseCaseSuite) TestDeleteTask_Success() {
	taskID := "some-task-id"

	suite.repository.EXPECT().
		DeleteTask(taskID, suite.ctx).
		Return(nil).
		Times(1)

	err := suite.usecase.DeleteTask(taskID, suite.ctx)
	suite.Nil(err, "Expected no error while deleting the task")
}

func (suite *taskUseCaseSuite) TestDeleteTask_Failure() {
	taskID := "some-task-id"
	errResp := &domain.ErrorResponse{Message: "Deletion failed"}

	suite.repository.EXPECT().
		DeleteTask(taskID, suite.ctx).
		Return(errResp).
		Times(1)

	err := suite.usecase.DeleteTask(taskID, suite.ctx)
	suite.Equal(errResp, err, "Expected error response while deleting the task")
}

func (suite *taskUseCaseSuite) TestGetTaskByID_Success() {
	taskID := "some-task-id"
	expectedTask := domain.TaskDTO{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "Pending",
	}

	suite.repository.EXPECT().
		GetTaskByID(taskID, suite.ctx).
		Return(expectedTask, nil).
		Times(1)

	result, err := suite.usecase.GetTaskByID(taskID, suite.ctx)
	suite.Nil(err, "Expected no error while getting the task by ID")
	suite.Equal(expectedTask, result, "Expected task to match")
}

func (suite *taskUseCaseSuite) TestGetTaskByID_Failure() {
	taskID := "some-task-id"
	errResp := &domain.ErrorResponse{Message: "Task not found"}

	suite.repository.EXPECT().
		GetTaskByID(taskID, suite.ctx).
		Return(domain.TaskDTO{}, errResp).
		Times(1)

	result, err := suite.usecase.GetTaskByID(taskID, suite.ctx)
	suite.Equal(errResp, err, "Expected error response while getting the task by ID")
	suite.Equal(domain.TaskDTO{}, result, "Expected empty task DTO on error")
}

func (suite *taskUseCaseSuite) TestGetTasks_Success() {
	expectedTasks := []domain.TaskDTO{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now().Add(48 * time.Hour),
			Status:      "Completed",
		},
	}

	suite.repository.EXPECT().
		GetTasks(suite.ctx).
		Return(expectedTasks, nil).
		Times(1)

	result, err := suite.usecase.GetTasks(suite.ctx)
	suite.Nil(err, "Expected no error while getting tasks")
	suite.Equal(expectedTasks, result, "Expected tasks to match")
}

func (suite *taskUseCaseSuite) TestGetTasks_Failure() {
	errResp := &domain.ErrorResponse{Message: "Failed to get tasks"}

	suite.repository.EXPECT().
		GetTasks(suite.ctx).
		Return(nil, errResp).
		Times(1)

	result, err := suite.usecase.GetTasks(suite.ctx)
	suite.Equal(errResp, err, "Expected error response while getting tasks")
	suite.Nil(result, "Expected no tasks on error")
}

func (suite *taskUseCaseSuite) TestUpdateTask_Success() {
	taskID := "some-task-id"
	updatedTask := domain.TaskCreateDTO{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now().Add(72 * time.Hour),
		Status:      "Completed",
	}

	suite.repository.EXPECT().
		UpdateTask(taskID, updatedTask, suite.ctx).
		Return(nil).
		Times(1)

	err := suite.usecase.UpdateTask(taskID, updatedTask, suite.ctx)
	suite.Nil(err, "Expected no error while updating the task")
}

func (suite *taskUseCaseSuite) TestUpdateTask_Failure() {
	taskID := "some-task-id"
	updatedTask := domain.TaskCreateDTO{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now().Add(72 * time.Hour),
		Status:      "Completed",
	}
	errResp := &domain.ErrorResponse{Message: "Update failed"}

	suite.repository.EXPECT().
		UpdateTask(taskID, updatedTask, suite.ctx).
		Return(errResp).
		Times(1)

	err := suite.usecase.UpdateTask(taskID, updatedTask, suite.ctx)
	suite.Equal(errResp, err, "Expected error response while updating the task")
}

func TestTaskUseCase(t *testing.T) {
	suite.Run(t, new(taskUseCaseSuite))
}
