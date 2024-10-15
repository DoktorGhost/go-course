package address

import (
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"net/rpc"
	"time"
)

type GeoService struct {
	Client *rpc.Client
}

func (g *GeoService) AddressSearch(input entities.SearchRequest) (entities.Response, error) {
	start := time.Now()

	var addressSearchResp entities.Response

	err := g.Client.Call("GeoProvider.AddressSearch", input, &addressSearchResp)
	if err != nil {
		fmt.Println("Error calling AddressSearch:", err)
		return addressSearchResp, err
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("AddressSearch").Observe(duration)

	return addressSearchResp, nil
}

func (g *GeoService) GeoCode(input entities.GeocodeRequest) (entities.Response, error) {
	start := time.Now()

	var geoCodeResp entities.Response

	err := g.Client.Call("GeoProvider.GeoCode", input, &geoCodeResp)
	if err != nil {
		fmt.Println("Error calling GeoCode:", err)
		return geoCodeResp, err
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("GeoCode").Observe(duration)

	return geoCodeResp, nil
}
