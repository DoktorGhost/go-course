package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"httpapi/internal/address"
	"httpapi/internal/auth"
)

func SetupRoutes(geoService address.GeoUseCase, authUseCase *auth.AuthUseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authUseCase.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", handleSearch(geoService))
		r.Post("/api/address/geocode", handleGeocode(geoService))
	})

	r.Post("/api/login", handleLogin(authUseCase))
	r.Post("/api/register", handleRegister(authUseCase))

	// Настройка Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
