package usecase

import (
	"auth/internal/entities"
	"auth/internal/grpcauth/client"
	"auth/internal/services"
	"errors"
)

type UseCase struct {
	AuthService *services.AuthService
	AuthClient  *client.AuthClient
}

func NewUseCase(authService *services.AuthService, authClient *client.AuthClient) *UseCase {
	return &UseCase{AuthService: authService, AuthClient: authClient}
}

func (uc *UseCase) Register(userData entities.User) (string, error) {
	id, err := uc.AuthClient.AddUser(userData.Email, userData.Password)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (uc *UseCase) Login(userData entities.User) (string, error) {
	valid, err := uc.AuthClient.CheckUser(userData.Email, userData.Password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", errors.New("invalid user")
	}
	token, err := uc.AuthService.GenerateJWT(userData.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UseCase) CheckToken(token string) error {
	claims, err := uc.AuthService.ValidateJWT(token)
	if err != nil {
		return err
	}
	if claims.Valid() != nil {
		return errors.New("invalid token")
	}
	return nil
}
