package user_usecase

import (
	"errors"
	"fmt"
	"geoservice/internal/auth"
	"geoservice/internal/entities"
	"geoservice/internal/services/user_services"
	"github.com/go-chi/jwtauth"
)

type UsersUseCase struct {
	userService *user_services.UserService
	TokenAuth   *jwtauth.JWTAuth
}

func NewUsersUseCase(userService *user_services.UserService, secretKey string) *UsersUseCase {
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	return &UsersUseCase{userService: userService, TokenAuth: tokenAuth}
}

func (uc *UsersUseCase) AddUser(userData entities.UserType) (int, error) {
	// Проверка, существует ли пользователь с таким именем
	_, err := uc.userService.GetUser(userData.Username)
	if err == nil {
		return 0, fmt.Errorf("пользователь с таким Username уже существует")
	}

	// Хеширование пароля
	hash, err := auth.HashPassword(userData.Password)
	if err != nil {
		return 0, fmt.Errorf("ошибка хеширования пароля: %v", err)
	}

	userData.Password = hash

	// Создание пользователя
	id, err := uc.userService.CreateUser(userData)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	return id, nil
}

// Login проверяет логин и пароль пользователя.
func (uc *UsersUseCase) Login(login, password string) (string, error) {
	// Получаем сохранённый хеш пароля пользователя
	userData, err := uc.userService.GetUser(login)
	if err != nil {
		return "", fmt.Errorf("пользователь с таким Username не найден: %v", err)
	}

	// Сравниваем пароль с хешем
	flag, err := auth.CheckPasswordHash(password, userData.Password)
	if err != nil {
		return "", fmt.Errorf("ошибка аутентификации: %v", err)
	}

	if !flag {
		return "", errors.New("invalid password")
	}

	jwt, err := auth.GenerateJWT(login, uc.TokenAuth)
	if err != nil {
		return "", fmt.Errorf("ошибка jwt: %v", err)
	}

	// Авторизация успешна, возвращаем JWT
	return jwt, nil
}
