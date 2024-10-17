package grpcProvider

import (
	"context"
	"strconv"
	"user/internal/entities"
	"user/internal/usecase"
	"user/pkg/proto"
)

type UserGRPCServer struct {
	uc *usecase.UseCase
	proto.UnimplementedUserServiceServer
}

func NewUserGRPCServer(uc *usecase.UseCase) *UserGRPCServer {
	return &UserGRPCServer{uc: uc}
}

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *proto.User) (*proto.CreateUserResponse, error) {
	user := entities.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Вызов метода юзкейса
	id, err := s.uc.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{Id: strconv.Itoa(id)}, nil
}

func (s *UserGRPCServer) GetUser(ctx context.Context, req *proto.User) (*proto.GetUserResponse, error) {
	err := s.uc.Login(req.Email, req.Password)
	if err != nil {
		return &proto.GetUserResponse{Success: false}, err
	}

	return &proto.GetUserResponse{Success: true}, nil
}
