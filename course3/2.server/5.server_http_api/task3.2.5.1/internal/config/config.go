package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKeyValue    string
	SecretKeyValue string
	SecretKeyJWT   string
}

// LoadConfig загружает конфигурацию из .env файла и возвращает структуру Config
func LoadConfig() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Чтение переменных окружения
	cfg := &Config{
		ApiKeyValue:    os.Getenv("API_KEY"),
		SecretKeyValue: os.Getenv("SECRET_KEY"),
		SecretKeyJWT:   os.Getenv("SECRET_KEY_JWT"),
	}

	return cfg, nil
}
