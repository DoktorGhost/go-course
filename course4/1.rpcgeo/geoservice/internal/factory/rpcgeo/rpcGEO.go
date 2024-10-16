package rpcgeo

import (
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"log"
	"net/rpc"
	"time"
)

type GeoServiceRpc struct {
	Client *rpc.Client
}

func NewGeoService() *GeoServiceRpc {
	client, err := rpc.Dial("tcp", "newgeoapp:1234")
	if err != nil {
		log.Println("Error connecting to RPC server:", err)
		return nil
	}
	return &GeoServiceRpc{client}
}

func (g *GeoServiceRpc) AddressSearch(input entities.SearchRequest) (entities.Response, error) {
	start := time.Now()

	var addressSearchResp entities.Response

	err := g.Client.Call("GeoProvider.AddressSearch", input, &addressSearchResp)
	if err != nil {
		log.Println("Error calling AddressSearch:", err)
		return addressSearchResp, err
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("AddressSearch").Observe(duration)

	return addressSearchResp, nil
}

func (g *GeoServiceRpc) GeoCode(input entities.GeocodeRequest) (entities.Response, error) {
	start := time.Now()

	var geoCodeResp entities.Response

	err := g.Client.Call("GeoProvider.GeoCode", input, &geoCodeResp)
	if err != nil {
		log.Println("Error calling GeoCode:", err)
		return geoCodeResp, err
	}

	duration := time.Since(start).Seconds()
	metrics.ApiDuration.WithLabelValues("GeoCode").Observe(duration)

	return geoCodeResp, nil
}
