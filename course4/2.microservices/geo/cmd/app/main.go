package main

import (
	"geo/internal/config"
	client "geo/internal/grpcgeo"
	"geo/internal/handlers"
	"geo/internal/service/cache"
	"geo/internal/service/dbservice"
	"geo/internal/service/geoservice"
	"geo/internal/storage/postgres"
	"geo/internal/usecase"
	"geo/pkg/storage/psg"
	redisdb "geo/pkg/storage/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"

	"os"
	"time"
)

func main() {
	//считываем переменные окружения
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Println(err)
	}

	//экземпляр бд psg
	var db *pgxpool.Pool
	for i := 0; i < 5; i++ {
		db, err = psg.InitStorage(&cfg.DbConfigPsg, schema)
		if err == nil {
			log.Println("Connected to database")
			break
		}
		log.Println("Ошибка подключения к бд:", err, "попытка ", i+1)
		time.Sleep(5 * time.Second)
	}

	//redis
	redisClient, err := redisdb.Init(&cfg.DbConfigRedis)
	if err != nil {
		log.Fatalf("Ошибка инициализации Redis: %v", err)
	}
	defer redisClient.Close()

	//инициализация клиента
	authClient, conn := client.InitAuthClient(&cfg.AuthConfig)
	defer conn.Close()

	geoRepo := postgres.NewGeoRepository(db)

	dbService := dbservice.NewDBService(geoRepo)
	caheService := cache.NewCacheService(redisClient)
	geoService := geoservice.NewGeoService(cfg.APIConfig.APIKey, cfg.APIConfig.SecretKey)

	uc := usecase.NewGeoUseCase(dbService, caheService, geoService)

	router := handlers.SetupRoutes(uc, authClient)

	log.Println("HTTP server listening on :", cfg.HttpProvider.HttpProviderPort)
	if err := http.ListenAndServe(":"+cfg.HttpProvider.HttpProviderPort, router); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}

}
