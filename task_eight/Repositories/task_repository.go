package repositories

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	db         *mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		db:         db,
		collection: collection,
	}
}

func (t *taskRepository) getCollection() *mongo.Collection {
	return t.db.Collection(t.collection)
}

func (t *taskRepository) GetTasks(ctx context.Context) ([]domain.TaskDTO, *domain.ErrorResponse) {
	collection := t.getCollection()
	var tasks []domain.TaskDTO

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, domain.InternalServerError("Error fetching tasks: " + err.Error())
	}

	err = cur.All(ctx, &tasks)
	if err != nil {
		return nil, domain.InternalServerError("Error decoding tasks: " + err.Error())
	}

	return tasks, nil
}

func (t *taskRepository) GetTaskByID(taskID string, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {
	collection := t.getCollection()
	var task domain.TaskDTO

	objectID, nErr := primitive.ObjectIDFromHex(taskID)

	if nErr != nil {
		return domain.TaskDTO{}, domain.InternalServerError("Error converting the given ID")
	}

	filter := bson.M{"_id": objectID}
	err := collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.TaskDTO{}, domain.NotFound("Task with the given ID not found")
		}
		return domain.TaskDTO{}, domain.InternalServerError("Error fetching task: " + err.Error())
	}

	return task, nil
}

func (t *taskRepository) CreateTask(newTask domain.TaskCreateDTO, ctx context.Context) (domain.TaskDTO, *domain.ErrorResponse) {
	collection := t.getCollection()

	inserRes, err := collection.InsertOne(ctx, newTask)
	if err != nil {
		return domain.TaskDTO{}, domain.InternalServerError("Internal server error while inserting the document: " + err.Error())
	}

	objectID, ok := inserRes.InsertedID.(primitive.ObjectID)
	if !ok {
		return domain.TaskDTO{}, domain.InternalServerError("Error converting the inserted ID")
	}

	task, nErr := t.GetTaskByID(objectID.Hex(), ctx)
	if nErr != nil {
		return domain.TaskDTO{}, nErr
	}

	return task, nil
}

func (t *taskRepository) DeleteTask(taskID string, ctx context.Context) *domain.ErrorResponse {
	collection := t.getCollection()

	_, err := t.GetTaskByID(taskID, ctx)
	if err != nil {
		return err
	}

	objectID, nErr := primitive.ObjectIDFromHex(taskID)

	if nErr != nil {
		return domain.InternalServerError("Error converting the given ID")
	}

	filter := bson.M{"_id": objectID}
	deleteRes, nErr := collection.DeleteOne(ctx, filter)
	if nErr != nil || deleteRes.DeletedCount == 0 {
		return domain.InternalServerError("Error deleting task: " + nErr.Error())
	}

	return nil
}

func (t *taskRepository) UpdateTask(taskID string, updatedTask domain.TaskCreateDTO, ctx context.Context) *domain.ErrorResponse {
	collection := t.getCollection()

	objectID, nErr := primitive.ObjectIDFromHex(taskID)

	if nErr != nil {
		return domain.InternalServerError("Error converting the given ID")
	}

	filter := bson.M{"_id": objectID}
	update := bson.D{{Key: "$set", Value: updatedTask}}

	updatedRes, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.InternalServerError("Error updating task: " + err.Error())
	}

	if updatedRes.MatchedCount == 0 {
		return domain.NotFound("Task not found for update")
	}

	if updatedRes.ModifiedCount == 0 {
		return domain.BadRequest("Task hasn't changed from last update")
	}

	return nil
}
