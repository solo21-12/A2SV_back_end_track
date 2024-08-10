package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
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
