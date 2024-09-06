package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	data := generateData(10)
	go func() {
		time.Sleep(1 * time.Second)
		close(data)
	}()
	for num := range data {
		fmt.Println(num)
	}
}

func generateData(n int) chan int {
	data := make(chan int, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randInt := rand.Intn(101)
			data <- randInt
		}()
	}

	go func() {
		wg.Wait()
		close(data)
	}()

	return data
}
