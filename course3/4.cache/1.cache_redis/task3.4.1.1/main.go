package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

// Структура для кеша
type cache struct {
	client *redis.Client
}

// Реализация метода Set для кеша
func (c *cache) Set(key string, value interface{}) error {
	// Сериализуем данные в JSON перед записью в Redis
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}
	// Устанавливаем значение с временем жизни ключа (например, 10 минут)
	return c.client.Set(key, data, 10*time.Minute).Err()
}

// Реализация метода Get для кеша
func (c *cache) Get(key string) (interface{}, error) {
	// Получаем значение из Redis
	data, err := c.client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("ключ %s не найден", key)
		}
		return nil, fmt.Errorf("ошибка получения данных: %v", err)
	}

	// Возвращаем JSON-строку (так как данные сериализованы)
	return data, nil
}

func NewCache(client *redis.Client) Cacher {
	return &cache{
		client: client,
	}
}

// Структура пользователя для примера
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Создание клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cache := NewCache(client)

	// Установка значения по ключу
	err := cache.Set("some:key", "value")
	if err != nil {
		fmt.Printf("Ошибка установки значения: %v\n", err)
		return
	}

	// Получение значения по ключу
	value, err := cache.Get("some:key")
	if err != nil {
		fmt.Printf("Ошибка получения значения: %v\n", err)
		return
	}
	fmt.Printf("Значение по ключу 'some:key': %v\n", value)

	// Создание структуры User
	user := &User{
		ID:   1,
		Name: "John",
		Age:  30,
	}

	// Установка структуры User по ключу
	err = cache.Set(fmt.Sprintf("user:%v", user.ID), user)
	if err != nil {
		fmt.Printf("Ошибка установки пользователя: %v\n", err)
		return
	}

	// Получение пользователя по ключу
	userData, err := cache.Get(fmt.Sprintf("user:%v", user.ID))
	if err != nil {
		fmt.Printf("Ошибка получения пользователя: %v\n", err)
		return
	}
	fmt.Printf("Данные по ключу 'user:1': %v\n", userData)

	// Десериализация пользователя обратно в структуру User
	var retrievedUser User
	err = json.Unmarshal([]byte(userData.(string)), &retrievedUser)
	if err != nil {
		fmt.Printf("Ошибка десериализации данных пользователя: %v\n", err)
		return
	}
	fmt.Printf("Пользователь: %+v\n", retrievedUser)
}

/*
Программа должна успешно устанавливать значение по ключу и получать его обратно.
Значение, полученное по ключу, должно соответствовать установленному значению.
Программа должна корректно обрабатывать ошибки, связанные с взаимодействием с Redis.
Все ответы и описания должны быть представлены на русском языке.
Компонент должен учитывать redis.Nil и возвращать ошибку not found by key %key%, если ключ не
найден.
*/
