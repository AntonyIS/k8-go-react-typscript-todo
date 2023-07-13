package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWS_DEFAULT_REGION    string
	AWS_ACCESS_KEY_ID     string
	AWS_ACCESS_SECRET_KEY string
	TableName             string
	Port                  string
}

func NewConfiguration() Config {

	var (
		serverPort            = os.Getenv("SERVER_PORT")
		AWS_DEFAULT_REGION    = os.Getenv("AWS_DEFAULT_REGION")
		AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
		AWS_ACCESS_SECRET_KEY = os.Getenv("AWS_ACCESS_SECRET_KEY")
		Tablename             = os.Getenv("TODO_TABLE")
	)

	return Config{
		Port:                  serverPort,
		TableName:             Tablename,
		AWS_DEFAULT_REGION:    AWS_DEFAULT_REGION,
		AWS_ACCESS_KEY_ID:     AWS_ACCESS_KEY_ID,
		AWS_ACCESS_SECRET_KEY: AWS_ACCESS_SECRET_KEY,
	}
}

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return nil
}
