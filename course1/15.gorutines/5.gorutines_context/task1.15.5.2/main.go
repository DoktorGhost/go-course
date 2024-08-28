package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var res string
	res = contextWithTimout(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
	res = contextWithTimout(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)
}

func contextWithTimout(ctx context.Context, contextTimout time.Duration, timeAfter time.Duration) string {
	ctx, cancel := context.WithTimeout(ctx, contextTimout)
	defer cancel()
	select {
	case <-ctx.Done():
		return "превышено время ожидания контекста"
	case <-time.After(timeAfter):
		return "превышено время ожидания"
	}
}
