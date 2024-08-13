package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
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
func (t *taskUseCase) GetTaskByID(taskID string, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {
	return t.taskRepository.GetTaskByID(taskID, ctx)

}
func (t *taskUseCase) CreateTask(newTask domain.TaskCreateDTO, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {
	return t.taskRepository.CreateTask(newTask, ctx)

}
func (t *taskUseCase) DeleteTask(taskID string, ctx context.Context) *domain.ErrorResponse {
	return t.taskRepository.DeleteTask(taskID, ctx)

}
func (t *taskUseCase) UpdateTask(taskID string, updatedTask domain.TaskCreateDTO, ctx context.Context) *domain.ErrorResponse {
	return t.taskRepository.UpdateTask(taskID, updatedTask, ctx)

}
