package repositories

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		db:         db,
		collection: collection,
	}
}

func (u *userRepository) GetRole(ctx context.Context) string {
	users, _ := u.getAllUsers(ctx)

	if len(users) == 0 {
		return "admin"
	}

	return "user"

}

func (u *userRepository) getCollection() *mongo.Collection {
	return u.db.Collection(u.collection)
}

func (u *userRepository) getAllUsers(ctx context.Context) ([]domain.UserDTO, *domain.ErrorResponse) {
	collection := u.getCollection()
	opt := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})

	cur, err := collection.Find(ctx, bson.D{}, opt)
	if err != nil {
		return nil, domain.InternalServerError("Internal server error")
	}

	var users []domain.UserDTO

	if err := cur.All(ctx, &users); err != nil {
		return nil, domain.InternalServerError("Error decoding users: " + err.Error())
	}

	return users, nil
}

func (u *userRepository) GetUserEmail(ctx context.Context, email string) (*domain.User, *domain.ErrorResponse) {
	// opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})
	collection := u.getCollection()

	filter := bson.D{{Key: "email", Value: email}}
	var user domain.User

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &domain.User{}, domain.NotFound("User with the given email not found")
		}

		return &domain.User{}, domain.InternalServerError("Error fetching user: " + err.Error())
	}

	return &user, nil
}
func (u *userRepository) CreateUser(ctx context.Context, user domain.UserCreateRequest) (domain.UserDTO, *domain.ErrorResponse) {
	collection := u.getCollection()

	// Check if the user already exists
	existingUser, err := u.GetUserEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		// User already exists
		return domain.UserDTO{}, domain.BadRequest("User already exists")
	}

	// Get user role
	role := u.GetRole(ctx)

	// Create new user document
	newUserCreated := domain.User{
		Email:    user.Email,
		Password: user.Password,
		Role:     role,
	}

	// Insert the new user
	insertRes, nErr := collection.InsertOne(ctx, newUserCreated)
	if nErr != nil {
		return domain.UserDTO{}, domain.InternalServerError("Something went wrong")
	}

	// Convert InsertedID to primitive.ObjectID
	objectID, ok := insertRes.InsertedID.(primitive.ObjectID)
	if !ok {
		return domain.UserDTO{}, domain.InternalServerError("Error while converting the ID")
	}

	// Fetch the newly created user
	newUser, eErr := u.GetUserID(ctx, objectID.Hex())
	if eErr != nil {
		return domain.UserDTO{}, domain.InternalServerError("Error fetching newly created user")
	}

	return newUser, nil
}
func (u *userRepository) GetUserID(ctx context.Context, id string) (domain.UserDTO, *domain.ErrorResponse) {
	var user domain.UserDTO
	collection := u.getCollection()

	objectID, cErr := primitive.ObjectIDFromHex(id)

	if cErr != nil {
		return domain.UserDTO{}, domain.InternalServerError("Error converting the id")
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})

	err := collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.UserDTO{}, domain.NotFound("User with the given ID not found")
		}

		return domain.UserDTO{}, domain.InternalServerError("Error fetching user: " + err.Error())
	}

	return user, nil
}
func (u *userRepository) PromoteUser(ctx context.Context, id string) *domain.ErrorResponse {
	collection := u.getCollection()
	user, err := u.GetUserID(ctx, id)

	if err != nil {
		return domain.NotFound("User Not found")
	}

	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{
		Key: "role", Value: "admin",
	}}}}

	updateRes, nErr := collection.UpdateOne(ctx, filter, update)

	if nErr != nil {
		return domain.InternalServerError("Error updating user: " + nErr.Error())
	}

	if updateRes.MatchedCount > 0 && updateRes.ModifiedCount == 0 {
		return domain.BadRequest("User already is an admin")
	}
	return nil
}
