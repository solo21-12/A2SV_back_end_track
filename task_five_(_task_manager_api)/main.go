package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"example.com/task_manager_api/data"
	"example.com/task_manager_api/model"
)

func getTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data.Tasks)
}

func GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range data.Tasks {
		if task.ID == id {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "The task with the given id not found"})
}

func postTask(ctx *gin.Context) {
	var newTask model.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		return
	}

	data.Tasks = append(data.Tasks, newTask)
	ctx.IndentedJSON(http.StatusCreated, newTask)
}

func deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			ctx.IndentedJSON(http.StatusNoContent, nil)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "The task with the given id not found"})

}
