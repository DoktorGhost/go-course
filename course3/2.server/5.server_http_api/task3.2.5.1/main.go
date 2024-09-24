package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"task3.2.5.1/internal/address"
	"task3.2.5.1/internal/handlers"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ApiKeyValue := os.Getenv("API_KEY")
	SecretKeyValue := os.Getenv("SECRET_KEY")

	//создаем экземпляр геосервиса
	geoService := address.NewGeoService(ApiKeyValue, SecretKeyValue)

	r := handlers.SetupRoutes(geoService)

	http.ListenAndServe(":8080", r)
}
