package main

import (
	"log"
	"net/http"
	"proxy/internal/config"
	"proxy/internal/delivery"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/api/user/", delivery.ProxyHandler("http://user:8080"))
	http.HandleFunc("/api/auth/", delivery.ProxyHandler("http://auth:8080"))
	http.HandleFunc("/api/address/", delivery.RateLimitMiddleware(delivery.ProxyHandler("http://geo:8080")))

	log.Println("Proxy server started on :8080")
	http.ListenAndServe(":8080", nil)
}
