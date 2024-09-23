package main

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"log"
	"net/http"
)

var logger *zap.SugaredLogger

func main() {
	r := chi.NewRouter()

	// Применение middleware для логирования с помощью zap logger

	rawLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer rawLogger.Sync()

	logger = rawLogger.Sugar()

	r.Use(LoggerMiddleware)

	// Здесь можно добавить ваши маршруты с различными методами

	r.Get("/1", handleRoute1)
	r.Post("/1", handleRoute2)
	r.Delete("/1", handleRoute3)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("Request received",
			"method", r.Method,
			"path", r.URL.Path,
		)
		next.ServeHTTP(w, r)
	})
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Get"))
}
func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Post"))
}
func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Delete"))
}
