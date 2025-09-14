package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string
	ServerPort string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	AppDB      string
}

func NewConfig() Config {
	// Loads env variables into enviroment
	godotenv.Load()

	return Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		AppDB:      os.Getenv("APP_DB"),
	}
}

var EnvConfig = NewConfig()
