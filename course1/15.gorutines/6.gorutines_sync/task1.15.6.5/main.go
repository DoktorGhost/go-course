package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	cache map[string]*User
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]*User),
	}
}

func (cache *Cache) Set(key string, user *User) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cache[key] = user
}

// /
func (cache *Cache) Get(key string) interface{} {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	user, ok := cache.cache[key]
	if !ok {
		return nil
	}
	return user
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
				ID:   i,
				Name: fmt.Sprintf("user-%d", i),
			})
		}(i)
	}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	wg.Wait()
}
