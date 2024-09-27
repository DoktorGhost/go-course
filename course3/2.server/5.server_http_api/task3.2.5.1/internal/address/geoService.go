package address

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"net/http"
	"net/url"
)

type GeoService struct {
	api       *suggest.Api
	apiKey    string
	secretKey string
}

func NewGeoService(apiKey, secretKey string) *GeoService {
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return nil
	}

	creds := client.Credentials{
		ApiKeyValue:    apiKey,
		SecretKeyValue: secretKey,
	}

	api := suggest.Api{
		Client: client.NewClient(endpointUrl, client.WithCredentialProvider(&creds)),
	}

	return &GeoService{
		api:       &api,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (g *GeoService) AddressSearch(input SearchRequest) (Response, error) {
	var result Response
	rawRes, err := g.api.Address(context.Background(), &suggest.RequestParams{Query: input.Query})
	if err != nil {
		return result, err
	}

	for _, r := range rawRes {
		if r.Data.City == "" || r.Data.Street == "" {
			continue
		}
		result.Addresses = append(result.Addresses, &Address{City: r.Data.City, Street: r.Data.Street, House: r.Data.House, Lat: r.Data.GeoLat, Lon: r.Data.GeoLon})
	}

	return result, nil
}

func (g *GeoService) GeoCode(input GeocodeRequest) (Response, error) {
	var result Response

	httpClient := &http.Client{}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return result, err
	}

	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(jsonData))
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", g.apiKey))

	resp, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	var geoCode GeoCode

	err = json.NewDecoder(resp.Body).Decode(&geoCode)
	if err != nil {
		return result, err
	}

	for _, r := range geoCode.Suggestions {
		var address Address
		address.City = r.Data.City
		address.Street = r.Data.Street
		address.House = r.Data.House
		address.Lat = r.Data.Lat
		address.Lon = r.Data.Lon

		result.Addresses = append(result.Addresses, &address)
	}

	return result, nil
}
