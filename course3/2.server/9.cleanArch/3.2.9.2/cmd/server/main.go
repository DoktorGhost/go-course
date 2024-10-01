package main

import (
	"httpapi/internal/address"
	"httpapi/internal/auth"
	"httpapi/internal/config"
	"httpapi/internal/handlers"
	"httpapi/internal/storage"
	"log"
	"net/http"
)

// @title Geo Service
// @version 0.1.0
// @description Гео-сервис

// @securityDefinitions.apikey BearerAuth
// @type apiKey
// @name Authorization
// @in header

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
	geoService := address.NewGeoUseCase(address.NewGeoService(cfg.ApiKeyValue, cfg.SecretKeyValue))

	r := handlers.SetupRoutes(*geoService, authUseCase)

	http.ListenAndServe(":8080", r)
}
