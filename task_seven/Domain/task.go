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
	GetTasks(ctx context.Context) ([]TaskDTO, error)
	GetTaskByID(taskID string, ctx context.Context) (TaskDTO, error)
	CreateTask(newTask TaskCreateDTO, ctx context.Context) (TaskDTO, error)
	DeleteTask(taskID string, ctx context.Context) error
	UpdateTask(taskID primitive.ObjectID, updatedTask TaskCreateDTO, ctx context.Context) error
}
