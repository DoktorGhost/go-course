package redis

import (
	"context"
	"errors"
	"geoservice/internal/config"
	"github.com/go-redis/redis/v8"
	"log"
)

func Init(conf *config.Config) (*redis.Client, error) {
	// Создаем контекст для работы с Redis
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.DB.Redis_host + ":" + conf.DB.Redis_port,
		Password: "",
		DB:       0,
	})

	// Проверка соединения с Redis
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("ошибка соединения с Redis: " + err.Error())
	}

	log.Println("Соединение с Redis успешно:", pong)
	return client, nil
}
