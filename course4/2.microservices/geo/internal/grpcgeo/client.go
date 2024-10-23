package client

import (
	"context"
	"geo/internal/config"
	pbUser "geo/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

type AuthClient struct {
	pbUser.AuthServiceClient
}

func InitAuthClient(cfg *config.AuthConfig) (*AuthClient, *grpc.ClientConn) {
	// Подключаемся к gRPC-серверу Auth
	conn, err := grpc.Dial(cfg.Auth_host+":"+cfg.Auth_port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Создаем gRPC-клиента для сервиса Auth
	authClient := pbUser.NewAuthServiceClient(conn)

	// Создаем сервис AUTH, который будет использовать этот клиент
	authService := &AuthClient{authClient}

	log.Println("Connected to Auth service port:", cfg.Auth_port)
	return authService, conn
}

func (a *AuthClient) ValidateToken(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод GetUser в сервисе USER
	resp, err := a.AuthServiceClient.ValidateToken(ctx, &pbUser.TokenRequest{Token: token})
	if err != nil {
		return false, err
	}

	return resp.GetValid(), nil
}
