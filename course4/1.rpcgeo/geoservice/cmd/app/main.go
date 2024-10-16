package main

import (
	"geoservice/internal/address"
	"geoservice/internal/cache"
	"geoservice/internal/config"
	"geoservice/internal/factory"
	"geoservice/internal/handlers"
	"geoservice/internal/metrics"
	"geoservice/internal/services/geo_services"
	"geoservice/internal/services/user_services"
	"geoservice/internal/storage/psg"
	"geoservice/internal/storage/redis"
	"geoservice/internal/usecase/geo_usecase"
	"geoservice/internal/usecase/user_usecase"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// @title Geo Service
// @version 0.1.0
// @description Гео-сервис

// @securityDefinitions.apikey BearerAuth
// @type apiKey
// @name Authorization
// @in header

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	//считываем переменные окружения
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Println(err)
	}

	//экземпляр бд
	store, err := psg.InitStorage(cfg.DB, schema)
	if err != nil {
		log.Printf("Error initializing storage: %v", err)
	}

	redisClient, err := redis.Init(cfg)
	if err != nil {
		log.Fatalf("Ошибка инициализации Redis: %v", err)
	}
	defer redisClient.Close()

	metrics.Init()

	ur := psg.NewUserRepository(store)
	gr := psg.NewGeoRepository(store)

	us := user_services.NewUserService(ur)
	gs := geo_services.NewGeorService(gr)

	uc := user_usecase.NewUsersUseCase(us, cfg.Secret.SecretKeyJWT)
	gu := geo_usecase.NewGeoUseCase(gs)

	sr := cache.SomeRepositoryImpl{gr}

	proxy := cache.NewSomeRepositoryProxy(&sr, redisClient)

	apiGeoService, err := factory.GetGeoServiceFactory(cfg.Protocol.Rpc_protocol)
	apiGeoUsecase := address.NewGeoUseCase(apiGeoService)

	r := handlers.SetupRoutes(apiGeoUsecase, gu, uc, proxy)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Запуск сервера на порту :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
