package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type SomeRepository interface {
	GetData() string
}
type SomeRepositoryImpl struct{}

func (r *SomeRepositoryImpl) GetData() string {
	// Здесь происходит запрос к базе данных
	fmt.Println("Идем в БД")
	time.Sleep(2 * time.Second)
	return "data"
}

type SomeRepositoryProxy struct {
	repository SomeRepository
	cache      *redis.Client
}

func NewSomeRepositoryProxy(repository SomeRepository, cache *redis.Client) *SomeRepositoryProxy {
	return &SomeRepositoryProxy{
		repository: repository,
		cache:      cache,
	}
}

func (r *SomeRepositoryProxy) GetData() string {
	// Здесь происходит проверка наличия данных в кэше
	cachedData, err := r.cache.Get("data_key").Result()
	if err == redis.Nil {
		// Если данных нет в кэше, то они запрашиваются у оригинального объекта и сохраняются в кэш
		data := r.repository.GetData()
		err = r.cache.Set("data_key", data, time.Hour).Err()
		if err != nil {
			log.Println("Ошибка сохранения в кэш:", err)
		}
		return data
	} else if err != nil {
		log.Println("Ошибка при получении данных из кэша", err)

	}

	// Если данные есть в кэше, то они возвращаются

	return cachedData
}

func main() {
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	repository := &SomeRepositoryImpl{}
	proxy := NewSomeRepositoryProxy(repository, cache)

	fmt.Println("Первый запрос")
	data1 := proxy.GetData()
	fmt.Println("Полученные данные", data1)

	fmt.Println("Второй запрос")
	data2 := proxy.GetData()
	fmt.Println("Полученные данные", data2)
}
