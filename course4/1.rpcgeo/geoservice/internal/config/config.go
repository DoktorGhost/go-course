package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	API    ApiConfig
	DB     DataBaseConfig
	Secret SecretConfig
}

type ApiConfig struct {
	ApiKeyValue    string `env:"API_KEY"`
	SecretKeyValue string `env:"SECRET_KEY"`
}

type DataBaseConfig struct {
	DB_host    string `env:"DB_HOST"`
	DB_port    string `env:"DB_PORT"`
	DB_name    string `env:"DB_NAME"`
	DB_login   string `env:"DB_LOGIN"`
	DB_pass    string `env:"DB_PASS"`
	Redis_host string `env:"REDIS_HOST"`
	Redis_port string `env:"REDIS_PORT"`
}

type SecretConfig struct {
	SecretKeyJWT string `env:"SECRET_KEY_JWT"`
}

// LoadConfig загружает конфигурацию из .env файла и возвращает структуру Config
func LoadConfig() (*Config, error) {
	config := &Config{}
	//считываем все переменны окружения в cfg
	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %v", err)
	}
	return config, nil
}
