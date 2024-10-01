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

	responder := NewResponder()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authUseCase.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", handleSearch(geoService, responder))
		r.Post("/api/address/geocode", handleGeocode(geoService, responder))
	})

	r.Post("/api/login", handleLogin(authUseCase, responder))
	r.Post("/api/register", handleRegister(authUseCase, responder))

	// Настройка Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
