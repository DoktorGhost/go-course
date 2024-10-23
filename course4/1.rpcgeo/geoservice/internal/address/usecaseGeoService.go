package address

import "geoservice/internal/entities"

type GeoUseCase struct {
	geoAPI GeoServiceRepository
}

func NewGeoUseCase(api GeoServiceRepository) *GeoUseCase {
	return &GeoUseCase{
		geoAPI: api,
	}
}

func (uc *GeoUseCase) SearchAddress(input entities.SearchRequest) (entities.Response, error) {
	return uc.geoAPI.AddressSearch(input)
}

func (uc *GeoUseCase) PerformGeocode(input entities.GeocodeRequest) (entities.Response, error) {
	return uc.geoAPI.GeoCode(input)
}
