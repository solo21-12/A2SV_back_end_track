package bootstrap

import (
	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET           string `mapstructure:"JWT_SECRET"`
	MONGO_URL            string `mapstructure:"MONGO_URL"`
	MONGO_DATABASE       string `mapstructure:"MONGO_DATABASE"`
	SERVER_ADDRESS       string `mapstructure:"SERVER_ADDRESS"`
	USER_COLLECTION      string `mapstructure:"USER_COLLECTION"`
	TASK_COLLECTION      string `mapstructure:"TASK_COLLECTION"`
	ALLOWED_USERS        string `mapstructure:"ALLOWED_USERS"`
	TEST_DATABASE        string `mapstructure:"TEST_DATABASE"`
	TEST_USER_COLLECTION string `mapstructure:"TEST_USER_COLLECTION"`
	TEST_TASK_COLLECTION string `mapstructure:"TEST_TASK_COLLECTION"`
}

func NewEnv() *Env {
	viper.AutomaticEnv()

	env := Env{}
	env.MONGO_URL = viper.GetString("MONGO_URL")
	env.MONGO_DATABASE = viper.GetString("MONGO_DATABASE")
	env.SERVER_ADDRESS = viper.GetString("SERVER_ADDRESS")
	env.JWT_SECRET = viper.GetString("JWT_SECRET")
	env.USER_COLLECTION = viper.GetString("USER_COLLECTION")
	env.TASK_COLLECTION = viper.GetString("TASK_COLLECTION")
	env.ALLOWED_USERS = viper.GetString("ALLOWED_USERS")
	env.TEST_DATABASE = viper.GetString("TEST_DATABASE")
	env.TEST_USER_COLLECTION = viper.GetString("TEST_USER_COLLECTION")
	env.TEST_TASK_COLLECTION = viper.GetString("TEST_TASK_COLLECTION")

	return &env
}
