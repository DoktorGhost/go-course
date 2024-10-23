package handlers

import (
	"encoding/json"
	"geo/internal/entities"
	client "geo/internal/grpcgeo"
	"geo/internal/usecase"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func SetupRoutes(uc *usecase.UseCase, authClient *client.AuthClient) *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware(authClient))

	r.Get("/api/address/search", handleSearch(uc))
	r.Get("/api/address/geocode", handleGeocode(uc))

	return r
}

func handleSearch(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entities.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

		if err != nil {
			log.Println("Bad request Decode", err)
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		address, err := uc.SearchAddress(req.Query)
		if err != nil {
			log.Println("Bad request SearchAddress", err)
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(address)

	}
}

func handleGeocode(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("StatusMethodNotAllowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entities.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println("Bad request Decode", err)
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := uc.PerformGeocode(req)
		if err != nil {
			log.Println("Bad request PerformGeocode", err)
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println("Успешный запрос")
		json.NewEncoder(w).Encode(resp)
	}
}
