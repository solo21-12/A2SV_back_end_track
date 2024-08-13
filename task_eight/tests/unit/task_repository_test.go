package unit_test

import (
	"context"
	"testing"
	"time"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	repositories "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Repositories"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositorySuite struct {
	DB             *mongo.Database
	Collection     *mongo.Collection
	ENV            *bootstrap.Env
	taskRepository domain.TaskRepository
	ctx            context.Context
	suite.Suite
}

func (suit *TaskRepositorySuite) SetupSuite() {
	env := bootstrap.NewEnv()
	suit.ENV = env
	client := bootstrap.NewMongoDatabase(env)

	suit.DB = client.Database(env.TEST_DATABASE)
	suit.Collection = suit.DB.Collection(env.TEST_TASK_COLLECTION)
	suit.ctx = context.Background()
	suit.taskRepository = repositories.NewTaskRepository(suit.DB, env.TEST_TASK_COLLECTION)
}

func (suit *TaskRepositorySuite) TearDownSuite() {
	err := suit.Collection.Drop(suit.ctx)
	suit.Require().NoError(err, "Error dropping the collection")

	err = suit.DB.Client().Disconnect(suit.ctx)
	suit.Require().NoError(err, "Error disconnecting from the database")
}

func (suit *TaskRepositorySuite) SetupTest() {
	err := suit.Collection.Drop(suit.ctx)
	suit.Require().NoError(err, "Error dropping the collection")
}

func (suite *TaskRepositorySuite) createTestTask() (domain.TaskDTO, *domain.ErrorResponse) {
	newTask := domain.TaskCreateDTO{
		Title:       "Test task",
		Description: "Test task description",
		DueDate:     time.Now().Add(time.Duration(time.Now().Day()) * 365),
		Status:      "Pending",
	}
	return suite.taskRepository.CreateTask(newTask, suite.ctx)
}

func (suit *TaskRepositorySuite) TestGetTasks() {
	_, err := suit.taskRepository.GetTasks(suit.ctx)

	if err != nil {
		suit.T().Errorf("Expected no error but got: %v", err.Message)
	}
}

func (suit *TaskRepositorySuite) TestCreateTask() {
	createdTask, err := suit.createTestTask()
	if err != nil {
		suit.T().Errorf("Received the following error while creating the task: %v", err.Message)
	}

	suit.Equal("Test task", createdTask.Title, "Created task and the provided task should be equal")
	suit.NotEmpty(createdTask, "Created task shouldn't be empty")
}

func (suit *TaskRepositorySuite) TestGetTaskByID() {

	createdTask, err := suit.createTestTask()

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while creating the task %v", err.Message)
	}

	retrievedTask, nErr := suit.taskRepository.GetTaskByID(createdTask.ID.Hex(), suit.ctx)

	if nErr != nil {
		suit.T().Errorf("Recieved the follwoing error while retrieving the task %v", nErr.Message)
	}

	suit.Equal(createdTask.Title, retrievedTask.Title, "tasks must be equal")
	suit.Equal(createdTask.ID, retrievedTask.ID, "IDs must be equal")
	suit.NotEmpty(retrievedTask, "Retrieved task shouldn't be empty")

}

func (suit *TaskRepositorySuite) TestDeleteTask_Positive() {

	createdTask, err := suit.createTestTask()

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while creating the task %v", err.Message)
	}

	err = suit.taskRepository.DeleteTask(createdTask.ID.Hex(), suit.ctx)

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while deleting the task %v", err.Message)
	}

	_, nErr := suit.taskRepository.GetTaskByID(createdTask.ID.Hex(), suit.ctx)

	suit.Error(nErr, "Error while retriving the deleted task")

}

func (suit *TaskRepositorySuite) TestDeleteTask_Negative() {

	createdTask, err := suit.createTestTask()

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while creating the task %v", err.Message)
	}

	err = suit.taskRepository.DeleteTask(createdTask.ID.Hex(), suit.ctx)

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while deleting the task %v", err.Message)
	}

	_, nErr := suit.taskRepository.GetTaskByID(createdTask.ID.Hex(), suit.ctx)

	if nErr == nil {
		suit.T().Errorf("Should have an error but isn't")
	}

}

func (suit *TaskRepositorySuite) TestUpdateTask() {

	createdTask, err := suit.createTestTask()

	if err != nil {
		suit.T().Errorf("Recieved the follwoing error while creating the task %v", err.Message)
	}

	updatedTask := domain.TaskCreateDTO{
		Title:       "Updated Task",
		Description: createdTask.Description,
		DueDate:     createdTask.DueDate,
		Status:      createdTask.Status,
	}

	err = suit.taskRepository.UpdateTask(createdTask.ID.Hex(), updatedTask, suit.ctx)

	if err != nil {
		suit.T().Errorf("Error while updating the task %v", err.Message)
	}

	retrievedTask, nErr := suit.taskRepository.GetTaskByID(createdTask.ID.Hex(), suit.ctx)

	if nErr != nil {
		suit.T().Errorf("Recieved the follwoing error while retrieving the task %v", nErr.Message)
	}

	suit.Equal(updatedTask.Title, retrievedTask.Title, "The updated value isn't perssistent")

}

func TestTaskRespoitorySuit(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
