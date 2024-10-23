package services

import (
	"fmt"
	"user/internal/entities"
)

type UserRepository interface {
	CreateUser(user entities.User) (int, error)
	GetUser(email string) (entities.User, error)
	GetAllUser() ([]entities.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.User) (int, error) {
	userID, err := s.repo.CreateUser(user)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) GetUser(username string) (entities.User, error) {
	user, err := s.repo.GetUser(username)

	if err != nil {
		return entities.User{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return user, nil
}

func (s *UserService) GetAllUser() ([]entities.User, error) {
	users, err := s.repo.GetAllUser()

	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователей: %v", err)
	}
	return users, nil
}
