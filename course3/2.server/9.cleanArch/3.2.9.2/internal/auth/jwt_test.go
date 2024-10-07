package auth

import (
	"testing"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Инициализация jwtAuth
	tokenAuth := jwtauth.New("HS256", []byte("secretKey"), nil)

	// Тест успешной генерации JWT
	token, err := GenerateJWT("testuser", tokenAuth)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Парсинг токена и проверка claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	assert.NoError(t, err)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "testuser", claims["username"])
	expirationTime := claims["exp"].(float64)
	assert.True(t, expirationTime > float64(time.Now().Unix()))
}

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	// Тест хеширования пароля
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Проверяем, что хеш отличается от исходного пароля
	assert.NotEqual(t, password, hashedPassword)

	// Повторное хеширование того же пароля должно дать другой хеш
	hashedPassword2, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEqual(t, hashedPassword, hashedPassword2)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"

	// Хешируем пароль
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)

	// Тест корректного сравнения пароля и хеша
	match, err := CheckPasswordHash(password, hashedPassword)
	assert.NoError(t, err)
	assert.True(t, match)

	// Тест некорректного пароля
	wrongPassword := "wrongpassword"
	match, err = CheckPasswordHash(wrongPassword, hashedPassword)
	assert.Error(t, err) // bcrypt возвращает ошибку при некорректном пароле
	assert.False(t, match)
}
