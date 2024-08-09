package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Role     string             `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string             `json:"email" binding:"required"`
	Role  string             `json:"role"`
}

