package entity

type SearchRequest struct {
	Query string `json:"query"`
}

type Response struct {
	Addresses []*Address `json:"addresses"`
}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lon"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
}

type Suggestion struct {
	Value             string  `json:"value"`
	UnrestrictedValue string  `json:"unrestricted_value"`
	Data              Address `json:"data"`
}

type GeoCode struct {
	Suggestions []Suggestion `json:"suggestions"`
}
