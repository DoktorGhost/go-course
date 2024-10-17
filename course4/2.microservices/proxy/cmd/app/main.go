package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Настраиваем обработчики для каждого сервиса
	http.HandleFunc("/api/user/", proxyHandler("http://user:8080"))
	http.HandleFunc("/api/auth/", proxyHandler("http://auth:8080"))
	http.HandleFunc("/api/address/", proxyHandler("http://geo:8080"))

	// Запускаем прокси-сервер на порту 8080
	log.Println("Proxy server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func proxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логируем исходный запрос
		log.Printf("Proxying request: %s %s", r.Method, r.URL.Path)

		// Парсим целевой URL
		proxyURL, err := url.Parse(target)
		if err != nil {
			log.Println("Failed to parse target URL:", err)
			http.Error(w, "Failed to parse target URL", http.StatusInternalServerError)
			return
		}

		// Обновляем адрес запроса (схема и хост)
		r.URL.Scheme = proxyURL.Scheme
		r.URL.Host = proxyURL.Host

		// Логируем новый адрес запроса
		log.Printf("Forwarding to: %s", r.URL.String())

		// Проксируем запрос к целевому сервису
		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			log.Println("Error forwarding request:", err)
			http.Error(w, "Error forwarding request", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Копируем заголовки ответа
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// Устанавливаем статус ответа
		w.WriteHeader(resp.StatusCode)

		// Копируем тело ответа
		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Println("Error copying response body:", err)
		}
	}
}
