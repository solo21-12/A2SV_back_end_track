package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"example.com/task_manager_api/model"
	"example.com/task_manager_api/data"
)

type TaskController struct {
	service *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

func (t *TaskController) GetTasksController(ctx *gin.Context) {

	tasks, err := t.service.GetTasks(context.TODO())

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByIDController(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := t.service.GetTaskByID(id, context.TODO())

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)

}

func (t *TaskController) PostTaskController(ctx *gin.Context) {
	var newTask model.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newTask.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	insertResult, err := t.service.CreateTask(newTask, context.TODO())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, ok := insertResult.InsertedID.(primitive.ObjectID)

	if !ok {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	createdTask, _ := t.service.GetTaskByID(objectID.Hex(), ctx)

	ctx.IndentedJSON(http.StatusCreated, gin.H{"created task": createdTask})
}

func (t *TaskController) DeleteTaskController(ctx *gin.Context) {
	id := ctx.Param("id")

	deletedResult, err := t.service.DeleteTask(id, context.TODO())

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if deletedResult.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)

}

func (t *TaskController) UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask model.Task

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format: " + err.Error()})
		return
	}

	if err := ctx.ShouldBind(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Use the request context
	updatedResult, err := t.service.UpdateTask(objectID, updatedTask, ctx)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the update result

	if updatedResult.MatchedCount > 0 && updatedResult.ModifiedCount > 0 {
		updatedTask.ID = id
		ctx.IndentedJSON(http.StatusOK, updatedTask)
	} else if updatedResult.MatchedCount > 0 && updatedResult.ModifiedCount == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "No update has been made to the file"})

	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found or not updated"})
	}
}

