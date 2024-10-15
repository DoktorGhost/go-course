package user_services

import (
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"time"
)

type UserRepository interface {
	CreateUser(user entities.UserType) (int, error)
	GetUser(username string) (entities.UserType, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.UserType) (int, error) {
	start := time.Now()

	userID, err := s.repo.CreateUser(user)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("CreateUser").Observe(duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) GetUser(username string) (entities.UserType, error) {
	start := time.Now()

	user, err := s.repo.GetUser(username)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("GetUser").Observe(duration)

	if err != nil {
		return entities.UserType{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return user, nil
}
