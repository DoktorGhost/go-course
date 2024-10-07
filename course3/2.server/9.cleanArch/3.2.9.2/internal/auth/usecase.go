package auth

import (
	"errors"
	"github.com/go-chi/jwtauth"
	"httpapi/internal/repository"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AuthUseCase struct {
	Repo      repository.StorageRepository
	TokenAuth *jwtauth.JWTAuth
}

func NewAuthUseCase(repo repository.StorageRepository, secretKey string) *AuthUseCase {
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	return &AuthUseCase{Repo: repo, TokenAuth: tokenAuth}
}

func (uc *AuthUseCase) Register(loginData Login) error {
	_, err := uc.Repo.Read(loginData.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Хешируем пароль
	hashedPassword, err := HashPassword(loginData.Password)
	if err != nil {
		return err
	}

	loginData.Password = hashedPassword

	// Сохраняем пользователя в хранилище
	return uc.Repo.Create(loginData.Username, loginData.Password)
}

// Login проверяет логин и пароль пользователя.
func (uc *AuthUseCase) Login(loginData Login) (string, error) {
	// Получаем сохранённый хеш пароля пользователя
	hashedPassword, err := uc.Repo.Read(loginData.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Сравниваем пароль с хешем
	flag, err := CheckPasswordHash(loginData.Password, hashedPassword)
	if err != nil {
		return "", err
	}

	if !flag {
		return "", errors.New("invalid password")
	}

	jwt, err := GenerateJWT(loginData.Username, uc.TokenAuth)

	// Авторизация успешна, возвращаем JWT
	return jwt, nil
}
