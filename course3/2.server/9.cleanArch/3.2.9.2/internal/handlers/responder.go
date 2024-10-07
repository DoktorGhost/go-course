package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
	ErrorUnauthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorForbidden(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct{}

func NewResponder() Responder {
	return &Respond{}
}
func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("responder json encode error:", err)
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	log.Println("http response bad request status code", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	response := Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("response writer error on write", err)
	}

}
func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	log.Println("http response forbidden", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	response := Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("response writer error on write", err)
	}
}

func (r *Respond) ErrorUnauthorized(w http.ResponseWriter, err error) {
	log.Println("http response Unauthorized", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	response := Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("response writer error on write", err)
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	if errors.Is(err, context.Canceled) {
		return
	}
	log.Println("http response internal error", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	response := Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("response writer error on write", err)
	}
}
