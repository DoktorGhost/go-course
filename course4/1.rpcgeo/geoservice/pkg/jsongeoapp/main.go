package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/url"
)

// Address структура для представления адреса
type Address struct {
	City   string
	Street string
	House  string
	Lat    string
	Lon    string
}

type Suggestion struct {
	Value             string  `json:"value"`
	UnrestrictedValue string  `json:"unrestricted_value"`
	Data              Address `json:"data"`
}

type GeoCode struct {
	Suggestions []Suggestion `json:"suggestions"`
}

// AddressSearchRequest запрос для поиска адреса
type AddressSearchRequest struct {
	Query string
}

// AddressSearchResponse ответ для поиска адреса
type AddressSearchResponse struct {
	Addresses []Address
}

// GeoCodeRequest запрос для геокодирования
type GeoCodeRequest struct {
	Lat string
	Lng string
}

// GeoCodeResponse ответ для геокодирования
type GeoCodeResponse struct {
	Addresses []Address
}

// GeoProvider структура для предоставления геокодирования
type GeoProvider struct {
	api       *suggest.Api
	apiKey    string
	secretKey string
}

// NewGeoProvider создает новый GeoProvider
func NewGeoProvider(apiKey, secretKey string) *GeoProvider {
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		log.Println("Error parsing endpoint url:", err)
		return nil
	}

	creds := client.Credentials{
		ApiKeyValue:    apiKey,
		SecretKeyValue: secretKey,
	}

	api := suggest.Api{
		Client: client.NewClient(endpointUrl, client.WithCredentialProvider(&creds)),
	}

	return &GeoProvider{
		api:       &api,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// AddressSearch метод для поиска адреса
func (g *GeoProvider) AddressSearch(input AddressSearchRequest, resp *AddressSearchResponse) error {
	log.Println("JSON-RPC. AddressSearch")
	rawRes, err := g.api.Address(context.Background(), &suggest.RequestParams{Query: input.Query})
	if err != nil {
		return err
	}

	for _, r := range rawRes {
		if r.Data.City == "" || r.Data.Street == "" {
			continue
		}
		resp.Addresses = append(resp.Addresses, Address{City: r.Data.City, Street: r.Data.Street, House: r.Data.House, Lat: r.Data.GeoLat, Lon: r.Data.GeoLon})
	}

	return nil
}

// GeoCode метод для геокодирования
func (g *GeoProvider) GeoCode(input GeoCodeRequest, resp *GeoCodeResponse) error {
	log.Println("JSON-RPC. GeoCode")

	httpClient := &http.Client{}

	jsonData, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", g.apiKey))

	respHTTP, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer respHTTP.Body.Close() // Закрываем resp.Body

	if respHTTP.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get valid response: %s", respHTTP.Status)
	}

	var geoCode GeoCode
	if err := json.NewDecoder(respHTTP.Body).Decode(&geoCode); err != nil {
		return err
	}

	for _, r := range geoCode.Suggestions {
		resp.Addresses = append(resp.Addresses, Address{
			City:   r.Data.City,
			Street: r.Data.Street,
			House:  r.Data.House,
			Lat:    r.Data.Lat,
			Lon:    r.Data.Lon,
		})
	}

	return nil
}

func main() {
	apiKey := "1d3b2a6c5330e1c6621cbf25ede1332a82df89bc"
	secretKey := "5453123054eafc1d0afe05464a95af40d34a23c1"

	gp := NewGeoProvider(apiKey, secretKey)
	if gp == nil {
		log.Println("Failed to create GeoProvider")
		return
	}

	// Регистрация GeoProvider как RPC-сервиса
	err := rpc.Register(gp)
	if err != nil {
		log.Println("Error registering GeoProvider:", err)
		return
	}

	// Создание TCP слушателя на порту 1235
	listener, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 1235...")
	// Обработка входящих соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		// Запуск нового горутины для обработки соединения
		go jsonrpc.ServeConn(conn)
	}
}
