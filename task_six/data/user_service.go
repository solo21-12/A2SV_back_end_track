package data

import (
	"context"
	"fmt"

	"example.com/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return string(hashedPassword), nil
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

		if err := cur.Decode(&curElem); err != nil {
			return nil, fmt.Errorf("error decoding user: %v", err)
		}

		users = append(users, curElem)

	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)

	}

	return users, nil
}

func (u *UserService) GetUser(ctx context.Context, email string) (model.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var curUser model.User

	err := u.collection.FindOne(ctx, filter).Decode(&curUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("no user find with the given email address: %s", email)

		}
		return model.User{}, fmt.Errorf("error retriving a user: %v", err)
	}

	return curUser, nil
}

func (u *UserService) CreateUser(ctx context.Context, user model.User) (*mongo.InsertOneResult, error) {
	_, err := u.GetUser(ctx, user.Email)

	if err == nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

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

func (u *UserService) GetUserID(ctx context.Context, id primitive.ObjectID) (model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var user model.User
	err := u.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("no user found with the give id %v", id)
		}

		return model.User{}, err
	}

	return user, nil
}

func (u *UserService) PromoteUser(ctx context.Context, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	user, err := u.GetUserID(ctx, id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key: "_id", Value: user.ID,
	}}
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{
				Key: "role", Value: "admin",
			},
		},
	}}

	updateResult, err := u.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, fmt.Errorf("error occured during updating the user  with email %v to admin", user.Email)
	}

	return updateResult, nil
}
