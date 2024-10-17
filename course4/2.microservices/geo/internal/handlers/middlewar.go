package handlers

import (
	client "geo/internal/grpcgeo"
	"net/http"
	"strings"
)

func AuthMiddleware(authClient *client.AuthClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем токен из заголовка Authorization
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			// Убираем приставку "Bearer " из заголовка Authorization, если она есть
			token = strings.TrimPrefix(token, "Bearer ")

			// Проверяем токен через сервис Auth
			valid, err := authClient.ValidateToken(token)
			if err != nil || !valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Если токен валиден, продолжаем выполнение запроса
			next.ServeHTTP(w, r)
		})
	}
}
