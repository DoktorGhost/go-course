package main

import (
	"fmt"
	"time"
)

func main() {
	timoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)

	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}

func timeout(timeout time.Duration) func() bool {

	res := make(chan struct{})

	go func() {
		defer close(res)
		time.Sleep(timeout)
	}()

	return func() bool {
		select {
		case <-res:
			return true
		default:
			return false
		}
	}
}
