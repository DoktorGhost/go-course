package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"student.vkusvill.ru/samsonov/go-course/course3/2.server/5.server_http_api/task3.2.5.1/address"
	"student.vkusvill.ru/samsonov/go-course/course3/2.server/5.server_http_api/task3.2.5.1/entity"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Чтение переменных окружения
	ApiKeyValue := os.Getenv("API_KEY")
	SecretKeyValue := os.Getenv("SECRET_KEY")

	//создаем экземпляр геосервиса
	geoService := address.NewGeoService(ApiKeyValue, SecretKeyValue)

	r := chi.NewRouter()

	r.Post("/api/address/search", handleSearch(geoService))
	r.Post("/api/address/geocode", handleGeocode(geoService))

	http.ListenAndServe(":8080", r)

	/*
		query1 := entity.SearchRequest{"Ростов орбитальная 74"}
		input := entity.GeocodeRequest{"55.878", "37.653"}

		result, err := geoService.AddressSearch(query1)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, s := range result.Addresses {
			fmt.Printf("%s, %s, %s, %s, %s\n", s.City, s.Street, s.House, s.Lat, s.Lon)
		}

		fmt.Println("ПЕРВЫЙ ВЫПОЛНЕН")

		result2, err := geoService.GeoCode(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, s := range result2.Addresses {
			fmt.Printf("%s, %s, %s, %s\n", s.City, s.Street, s.House, s.Lat)
		}

	*/

}

func handleSearch(geoService *address.GeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entity.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.AddressSearch(req)
		if err != nil {
			http.Error(w, "Service is not responding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func handleGeocode(geoService *address.GeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req entity.GeocodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := geoService.GeoCode(req)
		if err != nil {
			http.Error(w, "Service is not responding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
