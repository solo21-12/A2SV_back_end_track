package repositories

import (
	"context"
	"testing"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/stretchr/testify/suite"
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
	suit.userRepository = NewUserRepository(suit.DB, env.TEST_USER_COLLECTION)
}

func (suit *UserRepositorySuit) TearDownSuite() {

	err := suit.DB.Drop(suit.ctx)
	suit.Require().NoError(err, "Failed to drop the database")

	err = suit.DB.Client().Disconnect(suit.ctx)
	suit.Require().NoError(err, "Failed to disconnect from database")

}

func (suit *UserRepositorySuit) TestCreateUser_Positive() {
	user := domain.UserCreateRequest{
		Email:    "firstTest@gmail.com",
		Password: "strongpassword",
	}

	_, err := suit.userRepository.CreateUser(suit.ctx, user)

	suit.NoError(err, "No error when create user")
}

func (suit *UserRepositorySuit) TestCreateUser_NilPointer_Negative() {
	_, err := suit.userRepository.CreateUser(suit.ctx, domain.UserCreateRequest{})

	suit.Error(err, "create error while nil input returns erro")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuit))
}
