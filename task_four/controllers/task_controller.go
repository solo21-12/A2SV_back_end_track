package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"example.com/task_manager_api/model"
	"example.com/task_manager_api/services"
)

func GetTasksController(ctx *gin.Context) {

	tasks := services.GetTasks()
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByIDController(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := services.GetTaskByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)

}

func PostTaskController(ctx *gin.Context) {
	var newTask model.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newTask.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	err := services.CreateTask(newTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTaskController(ctx *gin.Context) {
	id := ctx.Param("id")

	err := services.DeleteTask(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, model.Task{})
}

func UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask model.Task

	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	err, updatedTask := services.UpdateTask(id, updatedTask)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedTask)

}
