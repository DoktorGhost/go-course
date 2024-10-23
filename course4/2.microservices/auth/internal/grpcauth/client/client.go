package client

import (
	"auth/internal/config"
	pbUser "auth/pkg/proto/proto_user"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type AuthClient struct {
	UserClient pbUser.UserServiceClient
}

func InitAuthClient(cfg *config.Config) (*AuthClient, *grpc.ClientConn) {
	// Подключаемся к gRPC-серверу USER
	conn, err := grpc.Dial(cfg.User_host+":"+cfg.User_port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Создаем gRPC-клиента для сервиса USER
	userClient := pbUser.NewUserServiceClient(conn)

	// Создаем сервис AUTH, который будет использовать этот клиент
	authService := &AuthClient{UserClient: userClient}

	log.Println("Connected to Auth service port:", cfg.User_port)
	return authService, conn
}

func (a *AuthClient) CheckUser(email, password string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод GetUser в сервисе USER
	resp, err := a.UserClient.GetUser(ctx, &pbUser.User{Email: email, Password: password})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (a *AuthClient) AddUser(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод CreateUser в сервисе USER
	resp, err := a.UserClient.CreateUser(ctx, &pbUser.User{Email: email, Password: password})
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
