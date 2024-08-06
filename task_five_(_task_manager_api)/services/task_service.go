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

func CreateTask(newTask model.Task) error {
	_, err := GetTaskByID(newTask.ID)

	if err == nil { // Task exists
		return fmt.Errorf("the task already exists")
	}

	// Add new task
	data.Tasks = append(data.Tasks, newTask)

	return nil
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
