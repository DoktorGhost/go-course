package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"time"
)

// HashMaper интерфейс с методами Set и Get
type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// HashMap структура с таблицей и функцией хэширования
type HashMap struct {
	table    []map[string]interface{}
	hashFunc func(string) int
	capacity int
}

// Конструктор новой хэш-карты с выбором хэш-функции через функциональные опции
func NewHashMap(capacity int, opts ...func(*HashMap)) *HashMap {
	hm := &HashMap{
		table:    make([]map[string]interface{}, capacity),
		capacity: capacity,
	}
	for _, opt := range opts {
		opt(hm)
	}
	return hm
}

// Хэш-функции для разных CRC
func WithHashCRC32() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = func(key string) int {
			hash := crc32.NewIEEE()
			hash.Write([]byte(key))
			return int(hash.Sum32() % uint32(hm.capacity))
		}
	}
}

func WithHashCRC64() func(*HashMap) {
	return func(hm *HashMap) {
		table := crc64.MakeTable(crc64.ECMA)
		hm.hashFunc = func(key string) int {
			hash := crc64.New(table)
			hash.Write([]byte(key))
			return int(hash.Sum64() % uint64(hm.capacity))
		}
	}
}

// Метод Set для установки значения по ключу
func (hm *HashMap) Set(key string, value interface{}) {
	if hm.hashFunc == nil {
		panic("hash function not set")
	}
	index := hm.hashFunc(key)
	if hm.table[index] == nil {
		hm.table[index] = make(map[string]interface{})
	}
	hm.table[index][key] = value
}

// Метод Get для получения значения по ключу
func (hm *HashMap) Get(key string) (interface{}, bool) {
	if hm.hashFunc == nil {
		panic("hash function not set")
	}
	index := hm.hashFunc(key)
	if hm.table[index] == nil {
		return nil, false
	}
	value, ok := hm.table[index][key]
	return value, ok
}

// Функция для измерения времени выполнения функции
func MeasureTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// Основная функция
func main() {
	hashFuncs := []struct {
		name     string
		hashFunc func(*HashMap)
	}{
		{"CRC32", WithHashCRC32()},
		{"CRC64", WithHashCRC64()},
	}

	for _, hf := range hashFuncs {
		fmt.Printf("Testing %s:\n", hf.name)
		m := NewHashMap(16, hf.hashFunc)
		since := MeasureTime(func() {
			m.Set("key", "value")
			if value, ok := m.Get("key"); ok {
				fmt.Println("Value:", value)
			}
		})
		fmt.Println("Time taken:", since)
	}
}
