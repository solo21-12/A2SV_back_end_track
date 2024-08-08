package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate *validator.Validate

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status" binding:"required"`
}

func init() {
	validate = validator.New()
}

func (t *Task) Validate() error {
	return validate.Struct(t)
}
