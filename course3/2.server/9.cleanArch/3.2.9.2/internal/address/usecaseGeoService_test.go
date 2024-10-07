package address

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockGeoAPI представляет собой мок реализации интерфейса GeoAPI
type MockGeoAPI struct {
	mock.Mock
}

func (m *MockGeoAPI) AddressSearch(input SearchRequest) (Response, error) {
	args := m.Called(input)
	return args.Get(0).(Response), args.Error(1)
}

func (m *MockGeoAPI) GeoCode(input GeocodeRequest) (Response, error) {
	args := m.Called(input)
	return args.Get(0).(Response), args.Error(1)
}

func TestGeoUseCase_SearchAddress(t *testing.T) {
	mockGeoAPI := new(MockGeoAPI)
	geoUseCase := NewGeoUseCase(mockGeoAPI)

	input := SearchRequest{Query: "Some address"}

	expectedResponse := Response{
		Addresses: []*Address{
			{City: "Test City", Street: "Test Street", House: "123", Lat: "55.7558", Lon: "37.6173"},
		},
	}

	mockGeoAPI.On("AddressSearch", input).Return(expectedResponse, nil)

	result, err := geoUseCase.SearchAddress(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)

	mockGeoAPI.AssertExpectations(t)
}

func TestGeoUseCase_PerformGeocode(t *testing.T) {
	mockGeoAPI := new(MockGeoAPI)
	geoUseCase := NewGeoUseCase(mockGeoAPI)

	input := GeocodeRequest{Lat: "55.7558", Lng: "37.6173"}

	expectedResponse := Response{
		Addresses: []*Address{
			{City: "Test City", Street: "Test Street", House: "123", Lat: "55.7558", Lon: "37.6173"},
		},
	}

	mockGeoAPI.On("GeoCode", input).Return(expectedResponse, nil)

	result, err := geoUseCase.PerformGeocode(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)

	mockGeoAPI.AssertExpectations(t)
}
