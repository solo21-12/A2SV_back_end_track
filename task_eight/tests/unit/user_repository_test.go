package unit_test

import (
	"context"
	"testing"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	repositories "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Repositories"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuit struct {
	suite.Suite
	DB             *mongo.Database
	Collection     *mongo.Collection
	ENV            *bootstrap.Env
	userRepository domain.UserRepository
	ctx            context.Context
}

func (suit *UserRepositorySuit) SetupSuite() {
	env := bootstrap.NewEnv()
	client := bootstrap.NewMongoDatabase(env)

	suit.DB = client.Database(env.TEST_DATABASE)
	suit.Collection = suit.DB.Collection(env.TEST_USER_COLLECTION)
	suit.ctx = context.Background()
	suit.ENV = env
	suit.userRepository = repositories.NewUserRepository(suit.DB, env.TEST_USER_COLLECTION)
}

func (suit *UserRepositorySuit) TearDownSuite() {

	err := suit.Collection.Drop(suit.ctx)
	suit.Require().NoError(err, "Failed to drop the collections")

	err = suit.DB.Client().Disconnect(suit.ctx)
	suit.Require().NoError(err, "Failed to disconnect from database")

}

func (suit *UserRepositorySuit) SetupTest() {
	err := suit.Collection.Drop(suit.ctx)
	suit.Require().NoError(err, "Failed to drop the collections")
}

func (suit *UserRepositorySuit) createTestUser(email string) (domain.UserDTO, *domain.ErrorResponse) {
	user := domain.UserCreateRequest{
		Email:    email,
		Password: "strongpassword",
	}

	return suit.userRepository.CreateUser(suit.ctx, user)
}

func (suit *UserRepositorySuit) TestCreateUser_Positive() {
	userEmail := "firstTest2@gmail.com"

	newUser, err := suit.createTestUser(userEmail)

	if err != nil {
		suit.T().Errorf("Expected no error but recived %v", err.Message)
	}
	suit.NotEmpty(newUser.Email, "Expected created user to have a non-empty email")
	suit.Equal(userEmail, newUser.Email, "Expected created user's email to match the provided email")
}

func (suit *UserRepositorySuit) TestCreateUser_NilPointer_Negative() {
	_, err := suit.userRepository.CreateUser(suit.ctx, domain.UserCreateRequest{})

	suit.Error(err, "create error while nil input returns erro")
}

func (suit *UserRepositorySuit) TestGetUserEmail_Positive() {

	userEmail := "firstTest2@gmail.com"

	newUser, err := suit.createTestUser(userEmail)

	if err != nil {
		suit.T().Errorf("Expected no error but got: %v", err)
	}

	user, nErr := suit.userRepository.GetUserEmail(suit.ctx, newUser.Email)

	if nErr != nil {
		suit.T().Errorf("Expected no error but got: %v", err)
	}

	suit.NotEmpty(user, "Expected returned user to have a non-empty email")
	suit.Equal(user.Email, userEmail, "Expected returned user's email to match the provided email")
}

func (suit *UserRepositorySuit) TestGetUserEmail_Negative() {
	email := "nouser@gmail.com"
	_, err := suit.userRepository.GetUserEmail(suit.ctx, email)
	suit.Error(err, "Expected error when user not found")
}

func (suit *UserRepositorySuit) TestGetUserID_Positive() {

	userEmail := "firstTest2@gmail.com"

	newUser, err := suit.createTestUser(userEmail)

	if err != nil {
		suit.T().Errorf("Expected no error but got: %v", err)
	}

	suit.NotEmpty(newUser.Email, "Expected created user to have a non-empty email")
	suit.Equal(userEmail, newUser.Email, "Expected created user's email to match the provided email")

	retrivedUser, rErr := suit.userRepository.GetUserID(suit.ctx, newUser.ID.Hex())

	if rErr != nil {
		suit.T().Errorf("Expected no error but got: %v", err)
	}

	suit.Equal(newUser.Email, retrivedUser.Email, "Expected retrieved user's email to match the created user's email")

}

func (suit *UserRepositorySuit) TestGetUserID_Negative() {

	id := primitive.NewObjectID()
	_, err := suit.userRepository.GetUserID(suit.ctx, id.Hex())

	suit.Error(err, "Expected error when user not found")
}

func (suit *UserRepositorySuit) TestPromoteUser_Positive() {
	adminEmail := "firstTest3@gmail.com"
	userEmail := "userTest4@gmail.com"

	_, err := suit.createTestUser(adminEmail)
	if err != nil {
		suit.T().Fatalf("Failed to create admin user: %v", err)
	}

	newUser, err := suit.createTestUser(userEmail)
	if err != nil {
		suit.T().Fatalf("Failed to create new user: %v", err)
	}

	unPromotedUser, err := suit.userRepository.GetUserID(suit.ctx, newUser.ID.Hex())
	if err != nil {
		suit.T().Fatalf("Failed to get user by ID before promotion: %v", err)
	}
	suit.Equal("user", unPromotedUser.Role, "User should be created as a user")

	err = suit.userRepository.PromoteUser(suit.ctx, newUser.ID.Hex())
	if err != nil {
		suit.T().Fatalf("Failed to promote user: %v", err)
	}

	retrivedUser, err := suit.userRepository.GetUserID(suit.ctx, newUser.ID.Hex())
	if err != nil {
		suit.T().Fatalf("Failed to get user by ID after promotion: %v", err)
	}
	suit.Equal(suit.ENV.ALLOWED_USERS, retrivedUser.Role, "The user should have been promoted to admin")
}

func (suit *UserRepositorySuit) TestPromoteUser_AlreadyAdmin() {
	adminEmail := "firstTest3@gmail.com"

	adminUser, err := suit.createTestUser(adminEmail)
	if err != nil {
		suit.T().Fatalf("Failed to create admin user: %v", err)
	}

	retrivedUser, err := suit.userRepository.GetUserID(suit.ctx, adminUser.ID.Hex())
	if err != nil {
		suit.T().Fatalf("Failed to get user by ID before promotion: %v", err)
	}
	suit.Equal(suit.ENV.ALLOWED_USERS, retrivedUser.Role, "The user should have been admin")

	err = suit.userRepository.PromoteUser(suit.ctx, retrivedUser.ID.Hex())
	suit.Error(err, "Should throw an error as the user is already an admin")

}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuit))
}
