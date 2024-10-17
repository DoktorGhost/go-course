package server

import (
	"auth/internal/config"
	"auth/internal/entities"
	"auth/internal/usecase"
	pbAuth "auth/pkg/proto/proto_auth"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pbAuth.UnimplementedAuthServiceServer
	UseCase usecase.UseCase
}

func InitAuthServer(cfg *config.Config, useCase usecase.UseCase) *Server {
	lis, err := net.Listen("tcp", ":"+cfg.Provider_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Регистрация сервера AUTH
	authServer := &Server{
		UseCase: useCase,
	}
	pbAuth.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("Starting gRPC server on port:", cfg.Provider_port)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return authServer // Возвращаем ссылку на инициализированный сервер
}

// Реализация метода Login
func (s *Server) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	user := entities.User{Email: req.Email, Password: req.Password}
	token, err := s.UseCase.Login(user)
	if err != nil {
		return &pbAuth.LoginResponse{Success: false}, err
	}
	return &pbAuth.LoginResponse{Token: token, Success: true}, nil
}

// Реализация метода ValidateToken
func (s *Server) ValidateToken(ctx context.Context, req *pbAuth.TokenRequest) (*pbAuth.TokenResponse, error) {
	err := s.UseCase.CheckToken(req.Token)
	if err != nil {
		return &pbAuth.TokenResponse{Valid: false}, err
	}

	return &pbAuth.TokenResponse{Valid: true}, nil
}
