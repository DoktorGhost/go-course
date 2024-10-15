package types

type Address struct {
	City   string
	Street string
	House  string
	Lat    string
	Lon    string
}

// AddressSearch запрос для поиска адреса
type AddressSearchRequest struct {
	Query string
}

// AddressSearch ответ для поиска адреса
type AddressSearchResponse struct {
	Addresses []Address
}

// GeoCode запрос для геокодирования
type GeoCodeRequest struct {
	Lat string
	Lng string
}

type Suggestion struct {
	Value             string  `json:"value"`
	UnrestrictedValue string  `json:"unrestricted_value"`
	Data              Address `json:"data"`
}

type GeoCode struct {
	Suggestions []Suggestion `json:"suggestions"`
}
