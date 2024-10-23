package handlers

import (
	"encoding/json"
	"errors"
	_ "geoservice/docs"
	"geoservice/internal/address"
	"geoservice/internal/cache"
	"geoservice/internal/entities"
	"geoservice/internal/usecase/geo_usecase"
	"log"
	"net/http"
)

// @Summary Поиск по адресу
// @Description Обрабатывает POST запросы для поиска адресов.
// @Tags гео-сервис
// @Accept json
// @Produce json
// @Param query body entities.SearchRequest true "Address Query"
// @Success 200 {object} entities.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/search [post]
// @Security BearerAuth
func handleSearch(apiGeoUseCase *address.GeoUseCase, dbGeoService *geo_usecase.GeoUseCase, responder Responder, proxy *cache.SomeRepositoryProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var req entities.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}

		//идем в кэш
		var resp entities.Response
		resp.Addresses, err = proxy.GetData(req.Query)
		if err == nil && len(resp.Addresses) > 0 {
			responder.OutputJSON(w, resp)
			return
		}

		resp, err = apiGeoUseCase.SearchAddress(req)
		log.Println("идем в АПИ")

		if err != nil {
			responder.ErrorInternal(w, err)
			return
		}
		err = dbGeoService.AddSearch(req.Query, resp.Addresses)
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
// @Param query body entities.GeocodeRequest true "Address Query"
// @Success 200 {object} entities.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/address/geocode [post]
// @Security BearerAuth
func handleGeocode(apiGeoUseCase *address.GeoUseCase, responder Responder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			responder.ErrorForbidden(w, errors.New("method not allowed"))
			return
		}

		var req entities.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responder.ErrorBadRequest(w, err)
			return
		}
		defer r.Body.Close()

		resp, err := apiGeoUseCase.PerformGeocode(req)
		if err != nil {
			responder.ErrorInternal(w, err)
			return
		}

		responder.OutputJSON(w, resp)
	}
}
