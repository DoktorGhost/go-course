package address

import "geoservice/internal/entities"

type GeoServiceRepository interface {
	AddressSearch(input entities.SearchRequest) (entities.Response, error)
	GeoCode(input entities.GeocodeRequest) (entities.Response, error)
}
