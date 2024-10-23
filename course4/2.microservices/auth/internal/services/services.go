package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthService struct {
	secretKey string
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{secretKey: secretKey}
}

// Генерация JWT токена
func (s *AuthService) GenerateJWT(email string) (string, error) {
	token, err := generateJWT(email, s.secretKey)
	if err != nil {
		return "", fmt.Errorf("error generating JWT: %v", err)
	}
	return token, nil
}

// Валидация JWT токена
func (s *AuthService) ValidateJWT(tokenString string) (jwt.Claims, error) {
	claims, err := validateJWT(tokenString, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("error validating JWT: %v", err)
	}
	return claims, nil
}

// Вспомогательная функция для генерации токена
func generateJWT(email, secretKey string) (string, error) {
	// Создаем токен с типом HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 10).Unix(), // Время истечения
	})

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

// Вспомогательная функция для валидации токена
func validateJWT(tokenString, secretKey string) (jwt.Claims, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Проверяем, валиден ли токен
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
