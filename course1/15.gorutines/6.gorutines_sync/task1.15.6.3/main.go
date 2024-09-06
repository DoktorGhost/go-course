package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.RWMutex
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) GetCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.Incr()
		}()

	}

	wg.Wait()

	fmt.Println("Count:", c.GetCount())
}
