package bootstrap

import (
	"log"
	"path/filepath"

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
	// This is to load the env file
	env := Env{}
	projectRoot, err := filepath.Abs(filepath.Join(".."))

	if err != nil {
		log.Fatalf("Error getting project root path: %v", err)
	}

	// Set the path to the .env file
	viper.SetConfigFile(filepath.Join(projectRoot, ".env"))

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Can't find the file .env: %v", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatalf("Environment can't be loaded : %v", err)
	}

	return &env
}
