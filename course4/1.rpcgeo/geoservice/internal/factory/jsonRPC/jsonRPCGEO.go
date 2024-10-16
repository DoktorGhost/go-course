package geojsonrpc

import (
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type GeoServiceJsonRpc struct {
	Client *rpc.Client
}

func NewGeoService() *GeoServiceJsonRpc {
	client, err := jsonrpc.Dial("tcp", "jsongeoapp:1235")
	if err != nil {
		log.Println("Error connecting to json-RPC server:", err)
		return nil
	}
	return &GeoServiceJsonRpc{client}
}

func (g *GeoServiceJsonRpc) AddressSearch(input entities.SearchRequest) (entities.Response, error) {
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

func (g *GeoServiceJsonRpc) GeoCode(input entities.GeocodeRequest) (entities.Response, error) {
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
