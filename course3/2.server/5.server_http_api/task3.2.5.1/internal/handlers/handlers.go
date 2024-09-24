package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "task3.2.5.1/docs"
	"task3.2.5.1/internal/address"
	"task3.2.5.1/internal/entity"
)

func SetupRoutes(geoService *address.GeoService) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/address/search", handleSearch(geoService))
	r.Post("/api/address/geocode", handleGeocode(geoService))

	// Настройка Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler) // Для Swagger UI
	r.Get("/swagger/index.html", httpSwagger.WrapHandler)

	return r
}

// @Summary Поиск по адресу
// @Description Обрабатывает POST запросы для поиска адресов.
// @Tags Поиск по адресу
// @Accept json
// @Produce json
// @Param query body entity.SearchRequest true "Address Query"
// @Success 200 {object} entity.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/search [post]
func handleSearch(geoService *address.GeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entity.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.AddressSearch(req)
		if err != nil {
			http.Error(w, "Service is not responding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

// @Summary Получение адреса по координатам
// @Description Обрабатывает POST запросы для получения адреса по координатам.
// @Tags Поиск по координатам
// @Accept json
// @Produce json
// @Param query body entity.GeocodeRequest true "Address Query"
// @Success 200 {object} entity.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/geocode [post]
func handleGeocode(geoService *address.GeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entity.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.GeoCode(req)
		if err != nil {
			http.Error(w, "Service is not responding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
