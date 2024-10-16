package factory

import (
	"fmt"
	"geoservice/internal/entities"
	geojsonrpc "geoservice/internal/factory/jsonRPC"
	georpc "geoservice/internal/factory/rpc"
)

type GeoServiceRepository interface {
	AddressSearch(input entities.SearchRequest) (entities.Response, error)
	GeoCode(input entities.GeocodeRequest) (entities.Response, error)
}

func GetGeoServiceFactory(protocol string) (GeoServiceRepository, error) {
	if protocol == "json-rpc" {
		return geojsonrpc.NewGeoService(), nil
	}

	if protocol == "rpc" {
		return georpc.NewGeoService(), nil
	}

	return nil, fmt.Errorf("No protocol")
}
