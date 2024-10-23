package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"geoservice/internal/storage/psg"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type SomeRepository interface {
	GetData(search string) ([]entities.Address, error)
}

type SomeRepositoryImpl struct {
	DB *psg.GeoRepository
}

func (r *SomeRepositoryImpl) GetData(search string) ([]entities.Address, error) {
	// Здесь запрос к базе данных
	addressess, err := r.DB.GetAddressesBySearchData(search)
	if err != nil {
		return nil, err
	}

	return addressess, nil
}

type SomeRepositoryProxy struct {
	repository SomeRepository
	cache      *redis.Client
}

// Конструктор для создания прокси
func NewSomeRepositoryProxy(repo SomeRepository, redisClient *redis.Client) *SomeRepositoryProxy {
	return &SomeRepositoryProxy{
		repository: repo,
		cache:      redisClient,
	}
}

// Метод для получения данных с использованием кэша
func (r *SomeRepositoryProxy) GetData(search string) ([]entities.Address, error) {
	// Попробуем сначала получить данные из кэша Redis

	start := time.Now()
	cachedData, err := r.cache.Get(context.Background(), search).Result()

	if err == redis.Nil {
		log.Println("в кэше нет данных")
		// Данных в кэше нет, делаем запрос к базе данных через оригинальный репозиторий
		data, err := r.repository.GetData(search)
		if err != nil {
			return nil, err
		}

		// Сериализуем данные в JSON перед сохранением в кэш
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("ошибка сериализации данных для кэша: %v", err)
		}

		// Сохраняем данные в кэш на некоторое время
		err = r.cache.Set(context.Background(), search, jsonData, 10*time.Minute).Err()
		if err != nil {
			return nil, fmt.Errorf("ошибка сохранения данных в кэш: %v", err)
		}

		return data, nil

	} else if err != nil {
		// Если ошибка произошла не из-за отсутствия данных в кэше
		return nil, fmt.Errorf("ошибка получения данных из кэша: %v", err)
	}

	var addresses []entities.Address
	err = json.Unmarshal([]byte(cachedData), &addresses)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования данных из кэша: %v", err)
	}

	duration := time.Since(start).Seconds()
	metrics.CacheDuration.WithLabelValues("GetData").Observe(duration)

	// Возвращаем данные из кэша
	return addresses, nil
}
