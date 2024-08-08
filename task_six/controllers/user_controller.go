package controller

import (
	"context"
	"net/http"

	"example.com/task_manager_api/model"
	"example.com/task_manager_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) RegisterUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// creating use in the database
	result, err := u.service.CreateUser(ctx, user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// type assertion for the object id
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Retirve user detail
	curUser, curErr := u.service.GetUserID(context.TODO(), objectId)
	if curErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": curErr.Error()})
		return
	}

	// Respond with the user registratio sucess message
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": model.UserRegister{
			ID:    objectId,
			Email: user.Email,
			Role:  curUser.Role,
		},
	})
}

func (u *UserController) LoginUser(ctx *gin.Context) {
	var user model.UserLogin

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	res, err := services.LoginUser(user, *u.service, ctx)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)
}
