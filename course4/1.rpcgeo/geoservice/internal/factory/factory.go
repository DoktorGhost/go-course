package factory

import (
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/factory/grpcgeo"
	"geoservice/internal/factory/jsonrpcgeo"
	"geoservice/internal/factory/rpcgeo"
)

type GeoServiceRepository interface {
	AddressSearch(input entities.SearchRequest) (entities.Response, error)
	GeoCode(input entities.GeocodeRequest) (entities.Response, error)
}

func GetGeoServiceFactory(protocol string) (GeoServiceRepository, error) {
	if protocol == "json-rpc" {
		return jsonrpcgeo.NewGeoService(), nil
	}
	if protocol == "rpc" {
		return rpcgeo.NewGeoService(), nil
	}
	if protocol == "grpc" {
		return grpcgeo.NewGeoServiceGrpc(), nil
	}

	return nil, fmt.Errorf("No protocol")
}
