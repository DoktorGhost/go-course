package main

import "fmt"

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(trySend(ch, 1))
		fmt.Println(trySend(ch, 2))
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
