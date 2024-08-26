package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	data := make(chan string)
	stop := make(chan struct{})

	go func() {
		time.Sleep(d)
		close(stop)
	}()

	go func() {
		defer close(data)
		for {
			select {
			case <-ticker.C:
				data <- message
			case <-stop:
				return
			}
		}

	}()

	return data

}
