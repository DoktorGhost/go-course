package handlers

import (
	"encoding/json"
	_ "httpapi/docs"
	"httpapi/internal/auth"
	"net/http"
)

// @Summary Логин
// @Description Логин пользователя и выдача JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param login body auth.Login true "User Login Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func handleLogin(uc *auth.AuthUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var loginData auth.Login
		err := json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Аутентификация пользователя
		token, err := uc.Login(loginData)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Успешная аутентификация — возвращаем токен
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary Регистрация
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body auth.Login true "User Registration Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/register [post]
func handleRegister(uc *auth.AuthUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var registerData auth.Login
		err := json.NewDecoder(r.Body).Decode(&registerData)
		if err != nil {
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Регистрация пользователя
		err = uc.Register(registerData)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Успешная регистрация
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"message": "Registration successful"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
