package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskrepository domain.TaskRepository) domain.TaskUseCase {
	return &taskUseCase{
		taskRepository: taskrepository,
	}
}

func (t *taskUseCase) GetTasks(ctx context.Context) ([]domain.TaskDTO, *domain.ErrorResponse) {
	return t.taskRepository.GetTasks(ctx)
}
func (t *taskUseCase) GetTaskByID(taskID primitive.ObjectID, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {

}
func (t *taskUseCase) CreateTask(newTask domain.TaskCreateDTO, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {

}
func (t *taskUseCase) DeleteTask(taskID primitive.ObjectID, ctx context.Context) *domain.ErrorResponse {

}
func (t *taskUseCase) UpdateTask(taskID primitive.ObjectID, updatedTask domain.TaskCreateDTO, ctx context.Context) *domain.ErrorResponse {

}
