package data

import (
	"context"
	"fmt"

	"example.com/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
	}
}

func (task *TaskService) GetTasks(ctx context.Context) ([]model.Task, error) {
	var allTasks []model.Task

	cur, err := task.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var curTask model.Task

		err := cur.Decode(&curTask)

		if err != nil {
			return nil, fmt.Errorf("error: %v", err.Error())
		}

		allTasks = append(allTasks, curTask)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return allTasks, nil
}

func (task *TaskService) GetTaskByID(taskID string, ctx context.Context) (model.Task, error) {
	var curTask model.Task

	objectID, e := primitive.ObjectIDFromHex(taskID)

	if e != nil {
		return model.Task{}, fmt.Errorf("invalid task ID format: %v", e)
	}

	// Create the filter using BSON ObjectId type if taskID is expected to be an ObjectId
	filter := bson.D{{Key: "_id", Value: objectID}}

	// Find the document in the collection
	err := task.collection.FindOne(ctx, filter).Decode(&curTask)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Task{}, fmt.Errorf("task with the given ID not found")
		}
		return model.Task{}, fmt.Errorf("error retrieving task: %v", err)
	}

	return curTask, nil
}

func (task *TaskService) CreateTask(newTask model.Task, ctx context.Context) (*mongo.InsertOneResult, error) {
	insertResult, err := task.collection.InsertOne(ctx, newTask)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (task *TaskService) DeleteTask(taskID string, ctx context.Context) (*mongo.DeleteResult, error) {
	objecctID, e := primitive.ObjectIDFromHex(taskID)

	if e != nil {
		return nil, fmt.Errorf("invalid task ID format: %v", e)
	}

	filter := bson.D{{Key: "_id", Value: objecctID}}
	deleteRes, err := task.collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return deleteRes, nil
}

func (task *TaskService) UpdateTask(taskID primitive.ObjectID, updatedTask model.Task, ctx context.Context) (*mongo.UpdateResult, error) {

	// Create the filter to find the document to update
	filter := bson.D{{Key: "_id", Value: taskID}}

	// Define the update operation using $set to update only the specified fields
	update := bson.D{{Key: "$set", Value: updatedTask}}

	// Perform the update operation
	updatedResult, err := task.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error updating task: %v", err)
	}

	return updatedResult, nil
}
