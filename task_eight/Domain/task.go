package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDTO struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status"`
}

type TaskCreateDTO struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status" binding:"required"`
}

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]TaskDTO, *ErrorResponse)
	GetTaskByID(taskID string, ctx context.Context) (TaskDTO, *ErrorResponse)
	CreateTask(newTask TaskCreateDTO, ctx context.Context) (TaskDTO, *ErrorResponse)
	DeleteTask(taskID string, ctx context.Context) *ErrorResponse
	UpdateTask(taskID string, updatedTask TaskCreateDTO, ctx context.Context) *ErrorResponse
}


type TaskUseCase interface {
	GetTasks(ctx context.Context) ([]TaskDTO, *ErrorResponse)
	GetTaskByID(taskID string, ctx context.Context) (TaskDTO, *ErrorResponse)
	CreateTask(newTask TaskCreateDTO, ctx context.Context) (TaskDTO, *ErrorResponse)
	DeleteTask(taskID string, ctx context.Context) *ErrorResponse
	UpdateTask(taskID string, updatedTask TaskCreateDTO, ctx context.Context) *ErrorResponse
}
