package services

import (
	"fmt"

	"example.com/task_manager_api/data"
	"example.com/task_manager_api/model"
)


func GetTasks() []model.Task {
	return data.Tasks
}

func GetTaskByID(taskID string) (model.Task, error) {
	for _, task := range data.Tasks {
		if task.ID == taskID {
			return task, nil
		}
	}

	return model.Task{}, fmt.Errorf("task with the given id not found")
}

func CreateTask(newTask model.Task) (error, model.Task) {
	// Generate a new ID based on the current length of tasks
	newTask.ID = fmt.Sprintf("%d", len(data.Tasks)+1)
	// Add new task
	data.Tasks = append(data.Tasks, newTask)

	return nil, newTask
}


func DeleteTask(taskID string) error {
	for i, task := range data.Tasks {
		if task.ID == taskID {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("the task with the given id not found")
}

func UpdateTask(taskID string, updatedTask model.Task) (error, model.Task) {

	for i, task := range data.Tasks {
		if task.ID == taskID {

			if updatedTask.Title != "" {
				data.Tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				data.Tasks[i].Description = updatedTask.Description
			}

			if updatedTask.Status != "" {
				data.Tasks[i].Status = updatedTask.Status
			}

			if !updatedTask.DueDate.IsZero() {
				data.Tasks[i].DueDate = updatedTask.DueDate
			}

			return nil, data.Tasks[i]
		}
	}

	return fmt.Errorf("task with the given id not found"), model.Task{}
}
