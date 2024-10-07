package handlers

import (
	"encoding/json"
	_ "httpapi/docs"
	"httpapi/internal/address"
	"net/http"
)

// @Summary Поиск по адресу
// @Description Обрабатывает POST запросы для поиска адресов.
// @Tags гео-сервис
// @Accept json
// @Produce json
// @Param query body address.SearchRequest true "Address Query"
// @Success 200 {object} address.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/search [post]
// @Security BearerAuth
func handleSearch(geoService address.GeoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req address.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.SearchAddress(req)
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
// @Tags гео-сервис
// @Accept json
// @Produce json
// @Param query body address.GeocodeRequest true "Address Query"
// @Success 200 {object} address.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/geocode [post]
// @Security BearerAuth
func handleGeocode(geoService address.GeoUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req address.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.PerformGeocode(req)
		if err != nil {
			http.Error(w, "Service is not responding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
