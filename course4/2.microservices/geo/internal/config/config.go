package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DbConfigPsg   DbConfigPsg
	AuthConfig    AuthConfig
	DbConfigRedis DbConfigRedis
	HttpProvider  HttpProvider
	APIConfig     APIConfig
}

type AuthConfig struct {
	Auth_port string `env:"AUTH_SERVICE_PORT"`
	Auth_host string `env:"AUTH_SERVICE_HOST"`
}

type DbConfigPsg struct {
	DB_login string `env:"DB_LOGIN"`
	DB_pass  string `env:"DB_PASS"`
	DB_host  string `env:"DB_HOST"`
	DB_port  string `env:"DB_PORT"`
	DB_name  string `env:"DB_NAME"`
}

type APIConfig struct {
	APIKey    string `env:"API_KEY"`
	SecretKey string `env:"API_SECRET"`
}

type HttpProvider struct {
	HttpProviderPort string `env:"HTTP_PROVIDER_PORT"`
}

type DbConfigRedis struct {
	Redis_host string `env:"REDIS_HOST"`
	Redis_port string `env:"REDIS_PORT"`
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
