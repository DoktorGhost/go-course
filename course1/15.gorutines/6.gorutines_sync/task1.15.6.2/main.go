package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() int {
	c.value++
	return c.value
}

func concurrentSafeCounter() int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter.Increment()
			fmt.Println(counter.value)
			mu.Unlock()
		}()
	}
	wg.Wait()
	return counter.value
}

func main() {

	i := concurrentSafeCounter()
	fmt.Printf("concurrentSafeCounter отработала %d раз\n", i)
}
