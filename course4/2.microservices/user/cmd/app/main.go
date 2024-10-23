package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	"user/internal/config"
	"user/internal/handlers"
	"user/internal/provider/grpcProvider"
	"user/internal/services"
	"user/internal/storage/postgres"
	"user/internal/usecase"
	"user/pkg/proto"
	"user/pkg/storage/psg"
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

	//экземпляр бд
	var db *pgxpool.Pool
	for i := 0; i < 5; i++ {
		db, err = psg.InitStorage(cfg, schema)
		if err == nil {
			log.Println("Connected to database")
			break
		}
		log.Println("Ошибка подключения к бд:", err, "попытка ", i+1)
		time.Sleep(5 * time.Second)
	}

	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userUsecase := usecase.NewUseCase(userService)

	router := handlers.SetupRoutes(userUsecase)

	lis, err := net.Listen("tcp", ":"+cfg.Provider_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userGRPCServer := grpcProvider.NewUserGRPCServer(userUsecase)

	proto.RegisterUserServiceServer(grpcServer, userGRPCServer)
	reflection.Register(grpcServer)

	go func() {
		log.Println("gRPC server listening on :", cfg.Provider_port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	log.Println("HTTP server listening on :", cfg.HttpProviderPort)
	if err := http.ListenAndServe(":"+cfg.HttpProviderPort, router); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
