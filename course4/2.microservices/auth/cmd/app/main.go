package main

import (
	"auth/internal/config"
	"auth/internal/grpcauth/client"
	"auth/internal/grpcauth/server"
	"auth/internal/handlers"
	"auth/internal/services"
	"auth/internal/usecase"
	"log"
	"net/http"
)

func main() {
	//считываем переменные окружения
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	//инициализация клиента
	authClient, conn := client.InitAuthClient(cfg)
	defer conn.Close()

	authService := services.NewAuthService(cfg.SecretKeyJWT)

	uc := usecase.NewUseCase(authService, authClient)

	//инициализация сервера
	server.InitAuthServer(cfg, *uc)

	router := handlers.SetupRoutes(uc)

	log.Println("HTTP server listening on :", cfg.HttpProviderPort)
	if err := http.ListenAndServe(":"+cfg.HttpProviderPort, router); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
