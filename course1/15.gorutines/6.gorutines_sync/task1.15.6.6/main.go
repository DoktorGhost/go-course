package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Cache struct {
	data sync.Map
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	v, ok := c.data.Load(key)
	return v, ok
}

func main() {
	var wg sync.WaitGroup

	cache := &Cache{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set("name"+strconv.Itoa(i), "jhon"+strconv.Itoa(i))
		}(i)
	}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, ok := cache.Get("name" + strconv.Itoa(i))
			if ok {
				fmt.Println(value)
			}

		}(i)
	}

	wg.Wait()
}
