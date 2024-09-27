package address

type GeoServiceRepository interface {
	AddressSearch(input SearchRequest) (Response, error)
	GeoCode(input GeocodeRequest) (Response, error)
}
