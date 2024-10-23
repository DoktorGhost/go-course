package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"user/internal/usecase"
)

func SetupRoutes(uc *usecase.UseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/user/profile", handleProfile(uc))
	r.Get("/api/user/list", handleList(uc))

	return r
}

func handleProfile(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodGet {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Извлекаем email из URL-параметров
		email := r.URL.Query().Get("email")
		if email == "" {
			log.Println("Bad request: email is required")
			http.Error(w, "Bad request: email is required", http.StatusBadRequest)
			return
		}

		user, err := uc.GetUser(email)
		if err != nil {
			log.Println("User not found", err)
			http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(user)
	}
}

func handleList(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodGet {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		users, err := uc.GetAllUser()
		if err != nil {
			log.Println("Users not found", err)
			http.Error(w, "Users not found: "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(users)
	}
}
