package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/1", handleRoute1)
	r.Post("/1", handleRoute2)
	r.Delete("/1", handleRoute3)

	http.ListenAndServe(":8080", r)
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
