package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	// Инициализация лимитера
	limiter := ratelimit.New(5, ratelimit.Per(time.Minute))

	startTime := time.Now()
	// Вызов 5 раз подряд
	for i := 1; i < 6; i++ {
		limiter.Take()
		fmt.Println("вызов №", i, time.Since(startTime))
	}

	// Шестой вызов
	limiter.Take()
	fmt.Println("вызов №", 6, time.Since(startTime))
}
