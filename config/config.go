package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PlaidClientID string
	PlaidSecret   string
	PlaidEnv      string
	DatabaseURL   string
	Port          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file %v", err)
		return nil, err
	}
	return &Config{
		PlaidClientID: os.Getenv("PLAID_CLIENT_ID"),
		PlaidSecret:   os.Getenv("PLAID_SECRET"),
		PlaidEnv:      os.Getenv("PLAID_ENV"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		Port:          os.Getenv("PORT"),
	}, nil
}
