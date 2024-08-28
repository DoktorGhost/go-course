package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
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

func (cache *Cache) Get(key string) *User {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	user, ok := cache.cache[key]
	if !ok {
		return nil
	}
	return user
}

func keyBuilder(keys ...string) string {
	var res string
	for i, key := range keys {
		if i < len(keys)-1 {
			res += key + ":"
		}
		res += key
	}
	return res
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
	time.Sleep(1 * time.Second)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}

	wg.Wait()
}
