package handlers

import (
	"encoding/json"
	"errors"
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
func handleSearch(geoService address.GeoUseCase, responder Responder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var req address.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.SearchAddress(req)
		if err != nil {
			responder.ErrorInternal(w, err)
			return
		}

		responder.OutputJSON(w, resp)
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
func handleGeocode(geoService address.GeoUseCase, responder Responder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var req address.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.PerformGeocode(req)
		if err != nil {
			responder.ErrorInternal(w, err)
			return
		}

		responder.OutputJSON(w, resp)
	}
}
