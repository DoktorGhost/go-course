package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(5 * time.Second)
	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}
}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		select {
		case <-timer.C:
			result <- "Таймер сработал"
		case <-stop:
			result <- "Горутина завершила работу раньше, чем таймер сработал"
		}
	}()

	return result
}
