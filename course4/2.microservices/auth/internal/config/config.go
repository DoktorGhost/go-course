package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	User_port        string `env:"USER_SERVICE_PORT"`
	User_host        string `env:"USER_SERVICE_HOST"`
	Provider_port    string `env:"PROVIDER_PORT"`
	SecretKeyJWT     string `env:"SECRET_KEY_JWT"`
	HttpProviderPort string `env:"HTTP_PROVIDER_PORT"`
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
