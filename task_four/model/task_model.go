package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

type Task struct {
	ID          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status" binding:"required"`
}

func init() {
	validate = validator.New()
}

func (t *Task) Validate() error {
	return validate.Struct(t)
}
