package handlers

import (
	"auth/internal/entities"
	"auth/internal/usecase"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
)

func SetupRoutes(uc *usecase.UseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/auth/register", handleRegister(uc))
	r.Post("/api/auth/login", handleLogin(uc))

	return r
}

func handleRegister(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user entities.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Bad request Decode", err)
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		id, err := uc.Register(user)
		if err != nil {
			log.Println("Internal server error", err)
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(id)
	}
}

func handleLogin(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Получаем токен из заголовка
		token := r.Header.Get("Authorization")
		if token != "" {
			// Извлекаем токен, если он в формате Bearer
			if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
				token = token[7:] // удаляем "Bearer " из токена
			}
			// Если токен есть, проверяем его
			err := uc.CheckToken(token)
			if err != nil {
				log.Println("Invalid token:", err)
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}
		} else {
			var user entities.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Println("Bad request:", err)
				http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			token, err = uc.Login(user)
			if err != nil {
				log.Println("Invalid email or password", err)
				http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
				return
			}
		}

		// Отправляем токен в ответе
		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
