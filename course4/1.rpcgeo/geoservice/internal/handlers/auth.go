package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "geoservice/docs"
	"geoservice/internal/entities"
	"geoservice/internal/usecase/user_usecase"
	"net/http"
)

// @Summary Логин
// @Description Логин пользователя и выдача JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param login body entities.UserType true "User Login Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func handleLogin(uc *user_usecase.UsersUseCase, responder Responder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var loginData entities.UserType
		err := json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}
		defer r.Body.Close()

		// Аутентификация пользователя
		token, err := uc.Login(loginData.Username, loginData.Password)
		if err != nil {
			responder.ErrorUnauthorized(w, err)
			return
		}

		// Успешная аутентификация — возвращаем токен
		response := map[string]string{"token": token}
		responder.OutputJSON(w, response)
	}
}

// @Summary Регистрация
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body entities.UserType true "User Registration Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/register [post]
func handleRegister(uc *user_usecase.UsersUseCase, responder Responder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var registerData entities.UserType
		err := json.NewDecoder(r.Body).Decode(&registerData)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}
		defer r.Body.Close()

		// Регистрация пользователя
		id, err := uc.AddUser(registerData)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}

		// Успешная регистрация
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"message": fmt.Sprintf("Registration successful. ID=%d", id)}
		responder.OutputJSON(w, response)
	}
}
