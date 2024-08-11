package bootstrap

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET      string `mapstructure:"JWT_SECRET"`
	MONGO_USER      string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD  string `mapstructure:"MONGO_PASSWORD"`
	MONGO_DATABASE  string `mapstructure:"MONGO_DATABASE"`
	SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
	USER_COLLECTION string `mapstructure:"USER_COLLECTION"`
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
