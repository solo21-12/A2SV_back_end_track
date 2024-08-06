package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "The task with the given id not found"})
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
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": err})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, model.Task{})
}
