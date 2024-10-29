package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var (
	once   sync.Once
	Config config
)

type config struct {
	Broker_message string `env:"MESSAGE_BROKER"`
	Rabbit_host    string `env:"RABBITMQ_HOST"`
	Rabbit_port    string `env:"RABBITMQ_PORT"`
	Rabbit_user    string `env:"RABBITMQ_USER"`
	Rabbit_pass    string `env:"RABBITMQ_PASS"`
	Broker_kafka   string `env:"KAFKA_BROKER"`
	Topic_kafka    string `env:"KAFKA_TOPIC"`
	Http_port      string `env:"HTTP_PORT"`
}

// LoadConfig загружает конфигурацию из .env файла и возвращает структуру Config
func LoadConfig() config {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Errorf("Error loading .env file")
		}
		err = env.Parse(&Config)
		if err != nil {
			log.Fatal("ошибка загрузки конфигурации: %v", err)
		}
		fmt.Println(Config)
	})
	return Config
}
