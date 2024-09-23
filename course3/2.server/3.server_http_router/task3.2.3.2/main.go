package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()

	r.Get("/group1/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleGroup(w, r, "Group 1")
	})
	r.Get("/group2/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleGroup(w, r, "Group 2")
	})
	r.Get("/group3/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleGroup(w, r, "Group 3")
	})

	http.ListenAndServe(":8080", r)
}

func handleGroup(w http.ResponseWriter, r *http.Request, groupName string) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id > 0 && id < 4 {
		w.WriteHeader(http.StatusOK)
		res := fmt.Sprintf("%s Привет, мир %d", groupName, id)
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
