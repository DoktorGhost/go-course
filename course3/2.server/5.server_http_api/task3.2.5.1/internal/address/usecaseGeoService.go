package address

type GeoUseCase struct {
	geoAPI GeoServiceRepository
}

func NewGeoUseCase(api GeoServiceRepository) *GeoUseCase {
	return &GeoUseCase{
		geoAPI: api,
	}
}

func (uc *GeoUseCase) SearchAddress(input SearchRequest) (Response, error) {
	return uc.geoAPI.AddressSearch(input)
}

func (uc *GeoUseCase) PerformGeocode(input GeocodeRequest) (Response, error) {
	return uc.geoAPI.GeoCode(input)
}
