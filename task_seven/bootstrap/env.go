package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	MONGO_URL  string `mapstructure:"MONGO_URL"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
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
