package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/1", getOneHandler)
	r.Get("/2", getTwoHandler)
	r.Get("/3", getThreeHandler)

	http.ListenAndServe(":8080", r)
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func getTwoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 2"))
}

func getThreeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 3"))
}
