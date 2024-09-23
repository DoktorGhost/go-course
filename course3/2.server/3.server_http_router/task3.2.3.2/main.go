package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()

	r.Get("/group1/{id}", getOneHandler)
	r.Get("/group2/{id}", getTwoHandler)
	r.Get("/group3/{id}", getThreeHandler)

	http.ListenAndServe(":8080", r)
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id > 0 && id < 4 {
		w.WriteHeader(http.StatusOK)
		res := fmt.Sprintf("Group 1 Привет, мир %d", id)
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func getTwoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id > 0 && id < 4 {
		w.WriteHeader(http.StatusOK)
		res := fmt.Sprintf("Group 2 Привет, мир %d", id)
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getThreeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id > 0 && id < 4 {
		w.WriteHeader(http.StatusOK)
		res := fmt.Sprintf("Group 3 Привет, мир %d", id)
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
