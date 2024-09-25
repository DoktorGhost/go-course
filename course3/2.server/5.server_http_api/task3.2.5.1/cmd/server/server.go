package main

import (
	"context"
	"httpapi/internal/address"
	"httpapi/internal/auth"
	"httpapi/internal/config"
	"httpapi/internal/handlers"
	"httpapi/internal/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	//считываем переменные окружения
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	//экземпляр бд
	store := storage.NewStorageMap()

	authUseCase := auth.NewAuthUseCase(store, cfg.SecretKeyJWT)

	//экземпляр геосервиса
	geoService := address.NewGeoService(cfg.ApiKeyValue, cfg.SecretKeyValue)

	r := handlers.SetupRoutes(geoService, authUseCase)

	// Создание HTTP-сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Создание канала для получения сигналов остановки
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	wg.Add(2)
	// Запуск сервера в отдельной горутине
	go func() {
		defer wg.Done()
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Запуск горутины для обработки сигналов
	go func() {
		defer wg.Done()
		sig := <-sigChan
		log.Println("Got signal:", sig)

		// Создание контекста с таймаутом для graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // Отмена контекста после завершения

		// Остановка сервера с использованием graceful shutdown
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	wg.Wait()
	log.Println("Server stopped gracefully")
}
