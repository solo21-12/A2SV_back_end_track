package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the MongoDB document structure for a user.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Role     string             `json:"role"`
}

// UserCreateRequest represents the data required to create a new user.
type UserCreateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserDTO represents the data structure for a user response.
type UserDTO struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string             `json:"email"`
	Role  string             `json:"role"`
}

// LoginResponse represents the response structure after a successful login.
type LoginResponse struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]UserDTO, error)
	GetUser(ctx context.Context, email string) (UserDTO, error)
	CreateUser(ctx context.Context, user UserCreateRequest) (UserDTO, error)
	GetUserID(ctx context.Context, id primitive.ObjectID) (UserDTO, error)
	PromoteUser(ctx context.Context, id primitive.ObjectID) error
}
