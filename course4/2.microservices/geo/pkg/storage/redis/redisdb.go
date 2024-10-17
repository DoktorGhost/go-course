package redisdb

import (
	"context"
	"errors"
	"geo/internal/config"
	"github.com/go-redis/redis/v8"
	"log"
)

func Init(conf *config.DbConfigRedis) (*redis.Client, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis_host + ":" + conf.Redis_port,
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
