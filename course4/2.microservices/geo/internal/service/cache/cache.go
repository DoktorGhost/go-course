package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geo/internal/entities"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheService struct {
	cache *redis.Client
}

func NewCacheService(redisClient *redis.Client) *CacheService {
	return &CacheService{
		cache: redisClient,
	}
}

// Метод для получения данных с использованием кэша
func (r *CacheService) GetData(search string) ([]entities.Address, error) {
	// Попробуем сначала получить данные из кэша Redis
	cachedData, err := r.cache.Get(context.Background(), search).Result()
	var addresses []entities.Address
	if err == redis.Nil {
		return addresses, errors.New("кэш пуст")
	} else if err != nil {
		// Если ошибка произошла не из-за отсутствия данных в кэше
		return addresses, fmt.Errorf("ошибка получения данных из кэша: %v", err)
	}

	err = json.Unmarshal([]byte(cachedData), &addresses)
	if err != nil {
		return addresses, fmt.Errorf("ошибка декодирования данных из кэша: %v", err)
	}

	// Возвращаем данные из кэша
	return addresses, nil
}

func (r *CacheService) SaveData(search string, addresses []entities.Address) error {
	// Сериализуем данные в JSON перед сохранением в кэш
	jsonData, err := json.Marshal(addresses)
	if err != nil {
		return fmt.Errorf("ошибка сериализации данных для кэша: %v", err)
	}

	// Сохраняем данные в кэш на некоторое время
	err = r.cache.Set(context.Background(), search, jsonData, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("ошибка сохранения данных в кэш: %v", err)
	}

	return nil
}
