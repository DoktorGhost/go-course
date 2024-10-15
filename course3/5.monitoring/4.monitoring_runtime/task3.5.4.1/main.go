package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

func monitorGoroutines(ctx context.Context, prevGoroutines int) {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Мониторинг завершен")
		case <-ticker.C:
			countGoroutines := runtime.NumGoroutine()
			delta := ((countGoroutines * 100) / prevGoroutines) - 100

			if delta > 20 {
				fmt.Println("⚠️ Предупреждение: Количество горутин увеличилось более чем на 20%!")
			} else if delta < -20 {
				fmt.Println("⚠️ Предупреждение: Количество горутин уменьшилось более чем на 20%!")
			}

			fmt.Printf("Текущее количество горутин: %d\n", countGoroutines)
			prevGoroutines = countGoroutines
		}

	}

}
func main() {
	g, ctx := errgroup.WithContext(context.Background())
	// Мониторинг горутин
	go func() {
		monitorGoroutines(ctx, runtime.NumGoroutine())
	}()
	// Имитация активной работы приложения с созданием горутин
	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}
	// Ожидание завершения всех горутин
	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Все горутины завершены. Программа завершена.")
}
