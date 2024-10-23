package handlers

import (
	"geoservice/internal/address"
	"geoservice/internal/cache"
	"geoservice/internal/metrics"
	"geoservice/internal/usecase/geo_usecase"
	"geoservice/internal/usecase/user_usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRoutes(apiGeoUseCase *address.GeoUseCase, dbGeoService *geo_usecase.GeoUseCase, authUseCase *user_usecase.UsersUseCase, proxy *cache.SomeRepositoryProxy) *chi.Mux {
	r := chi.NewRouter()

	responder := NewResponder()

	r.Use(metrics.MetricsMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authUseCase.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", handleSearch(apiGeoUseCase, dbGeoService, responder, proxy))
		r.Post("/api/address/geocode", handleGeocode(apiGeoUseCase, responder))

		r.Get("/debug/pprof/", PprofHandler)
	})

	r.Post("/api/login", handleLogin(authUseCase, responder))
	r.Post("/api/register", handleRegister(authUseCase, responder))

	// Настройка Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Эндпоинт для метрик
	r.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics.PrometheusHandler().ServeHTTP(w, r)
	})

	return r
}
