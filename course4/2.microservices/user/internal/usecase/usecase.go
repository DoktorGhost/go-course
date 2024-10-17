package usecase

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"user/internal/entities"
	"user/internal/services"
)

type UseCase struct {
	userService *services.UserService
}

func NewUseCase(userService *services.UserService) *UseCase {
	return &UseCase{userService: userService}
}

func (uc *UseCase) CreateUser(userData entities.User) (int, error) {
	// Проверка, существует ли пользователь с таким именем
	_, err := uc.userService.GetUser(userData.Email)
	if err == nil {
		return 0, fmt.Errorf("пользователь с таким email уже существует")
	}

	// Хеширование пароля
	hash, err := hashPassword(userData.Password)
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

// проверяем логин и пароль пользователя
func (uc *UseCase) Login(email, password string) error {
	// Получаем сохранённый хеш пароля пользователя
	userData, err := uc.userService.GetUser(email)
	if err != nil {
		return fmt.Errorf("пользователь с таким email не найден: %v", err)
	}

	// Сравниваем пароль с хешем
	flag, err := checkPasswordHash(password, userData.Password)
	if err != nil {
		return fmt.Errorf("ошибка аутентификации: %v", err)
	}

	if !flag {
		return errors.New("invalid password")
	}

	// Авторизация успешна
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Получение пользователя
func (uc *UseCase) GetUser(email string) (entities.User, error) {
	var user entities.User
	userData, err := uc.userService.GetUser(email)
	if err != nil {
		return user, fmt.Errorf("пользователь с таким email не найден: %v", err)
	}

	return userData, nil
}

// Получение пользователей
func (uc *UseCase) GetAllUser() ([]entities.User, error) {
	var users []entities.User
	usersData, err := uc.userService.GetAllUser()
	if err != nil {
		return users, fmt.Errorf("пользователь с таким email не найден: %v", err)
	}

	return usersData, nil
}
