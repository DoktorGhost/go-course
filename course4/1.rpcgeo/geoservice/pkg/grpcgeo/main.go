package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"google.golang.org/grpc"
	pb "grpcgeo/generated"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
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

type server struct {
	pb.UnimplementedGeoServiceServer
	provider *GeoProvider
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
func (s *server) AddressSearch(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println("gRPC. AddressSearch")

	rawRes, err := s.provider.api.Address(context.Background(), &suggest.RequestParams{Query: req.Query})
	if err != nil {
		return nil, err
	}

	addresses := []*pb.Address{}

	for _, r := range rawRes {
		if r.Data.City == "" || r.Data.Street == "" {
			continue
		}
		addresses = append(addresses, &pb.Address{
			City:   r.Data.City,
			Street: r.Data.Street,
			House:  r.Data.House,
			Lat:    r.Data.GeoLat,
			Lon:    r.Data.GeoLon,
		})
	}

	return &pb.SearchResponse{Addresses: addresses}, nil
}

// GeoCode метод для геокодирования
func (s *server) GeoCode(ctx context.Context, req *pb.GeoCodeRequest) (*pb.GeoCodeResponse, error) {
	log.Println("gRPC. GeoCode")

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Token %s", s.provider.apiKey))

	respHTTP, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respHTTP.Body.Close() // Закрываем resp.Body

	if respHTTP.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get valid response: %s", respHTTP.Status)
	}

	var geoCode GeoCode
	if err := json.NewDecoder(respHTTP.Body).Decode(&geoCode); err != nil {
		return nil, err
	}

	addresses := []*pb.Address{}
	for _, r := range geoCode.Suggestions {
		addresses = append(addresses, &pb.Address{
			City:   r.Data.City,
			Street: r.Data.Street,
			House:  r.Data.House,
			Lat:    req.Lat,
			Lon:    req.Lng,
		})
	}

	return &pb.GeoCodeResponse{Addresses: addresses}, nil
}

func main() {
	apiKey := "1d3b2a6c5330e1c6621cbf25ede1332a82df89bc"
	secretKey := "5453123054eafc1d0afe05464a95af40d34a23c1"

	gp := NewGeoProvider(apiKey, secretKey)
	if gp == nil {
		log.Println("Failed to create GeoProvider")
		return
	}

	lis, err := net.Listen("tcp", ":1236")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGeoServiceServer(grpcServer, &server{provider: gp})

	log.Println("Listening on port 1236...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
