package main

import (
	"fmt"
	"sync"
)

// функция для параллельного считывания данных из каналов, порядок передачи не гарантирован
func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup

	for _, ch := range from {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for num := range c {
				mergeTo <- num
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(mergeTo)
	}()
}

// функция для параллельного считывания данных из каналов, порядок передачи не гарантирован
func mergeChan2(chans ...chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for num := range c {
				res <- num
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	c1 := generateChan(5)
	c2 := generateChan(5)
	c3 := generateChan(5)
	c4 := generateChan(5)
	c5 := generateChan(5)

	merge1 := make(chan int)

	mergeChan(merge1, c1, c2, c3, c4, c5)

	for num := range merge1 {
		fmt.Println(num)
	}

	c6 := generateChan(5)
	c7 := generateChan(5)
	c8 := generateChan(5)
	c9 := generateChan(5)
	c10 := generateChan(5)

	merge2 := mergeChan2(c6, c7, c8, c9, c10)

	for num := range merge2 {
		fmt.Println(num)
	}
}
