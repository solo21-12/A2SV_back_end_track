package services

import (
	"context"
	"fmt"

	"example.com/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		collection: db.Collection("users"),
	}
}

func generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err

}

func (u *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	filter := bson.D{}

	cur, err := u.collection.Find(ctx, filter)

	if err != nil {
		return []model.User{}, fmt.Errorf("error occured")
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var curElem model.User

		err := cur.Decode(&curElem)

		if err != nil {
			return nil, fmt.Errorf("error: %v", err.Error())
		}

		users = append(users, curElem)

	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())

	}

	return users, nil
}

func (u *UserService) GetUser(ctx context.Context, email string) (model.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var curUser model.User

	err := u.collection.FindOne(ctx, filter).Decode(&curUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("no user find with the given email address")

		}
		return model.User{}, fmt.Errorf("error retriving a user")
	}

	return curUser, nil
}

func (u *UserService) CreateUser(ctx context.Context, user model.User) (*mongo.InsertOneResult, error) {

	users, err := u.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, pErr := generatePassword(user.Password)

	if pErr != nil {
		return nil, fmt.Errorf("internal server error")
	}

	user.Password = hashedPassword

	insertRes, err := u.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("error while inserting the documents %v", err.Error())
	}

	return insertRes, nil

}
