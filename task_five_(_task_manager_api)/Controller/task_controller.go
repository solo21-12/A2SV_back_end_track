package services

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
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "The task with the given id not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)

}

func PostTaskController(ctx *gin.Context) {
	var newTask model.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		return
	}

	exist := services.CreateTask(newTask)

	if exist != nil {
		ctx.IndentedJSON(http.StatusBadRequest, exist)
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
