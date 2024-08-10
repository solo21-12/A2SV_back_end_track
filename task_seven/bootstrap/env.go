package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET     string `mapstructure:"JWT_SECRET"`
	MONGO_USER     string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD string `mapstructure:"MONGO_PASSWORD"`
	MONGO_DATABASE string `mapstructure:"MONGO_DATABASE"`
}

func NewEnv() *Env {
	// This is to load the env file
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Can't find the file .env: %v", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatalf("Environment can't be loaded : %v", err)
	}

	return &env
}
