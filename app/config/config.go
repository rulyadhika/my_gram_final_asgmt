package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	SERVER_PORT string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		DB_DRIVER:   os.Getenv("DB_DRIVER"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}
}
