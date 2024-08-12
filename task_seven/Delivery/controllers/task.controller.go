package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUseCase domain.TaskUseCase
}


func (t *TaskController) convertID(id string) (primitive.ObjectID, error) {

	return primitive.ObjectIDFromHex(id)
}

func (t *TaskController) Create(ctx *gin.Context) {
	var task domain.TaskCreateDTO

	if err := ctx.ShouldBind(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	newTask, err := t.TaskUseCase.CreateTask(task, ctx)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)

}

func (t *TaskController) GetAll(ctx *gin.Context) {
	tasks, err := t.TaskUseCase.GetTasks(ctx)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (t *TaskController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	objectID, err := t.convertID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format" + err.Error()})
		return
	}

	task, gErr := t.TaskUseCase.GetTaskByID(objectID, ctx)

	if gErr != nil {
		ctx.IndentedJSON(gErr.Code, gin.H{"error": gErr.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask domain.TaskCreateDTO

	objectID, oErr := t.convertID(id)
	if oErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid object id" + oErr.Error()})
		return
	}

	if err := ctx.ShouldBind(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data format" + err.Error()})
		return
	}

	err := t.TaskUseCase.UpdateTask(objectID, updatedTask, ctx)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	task, _ := t.TaskUseCase.GetTaskByID(objectID, ctx)

	ctx.IndentedJSON(http.StatusOK, task)

}

func (t *TaskController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	objectID, oErr := t.convertID(id)
	if oErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid object id" + oErr.Error()})
		return
	}

	if err := t.TaskUseCase.DeleteTask(objectID, ctx); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "task deleted successfully"})
}
