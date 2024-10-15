package geo_services

import (
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/metrics"
	"time"
)

type GeoRepository interface {
	AddSearchHistory(search string) (int, error)
	AddAddress(data entities.Address) (int, error)
	AddHistorySearchAddress(id_search, id_address int) (int, error)
	GetAddressesBySearchData(search string) ([]entities.Address, error)
}

type GeoService struct {
	repo GeoRepository
}

func NewGeorService(repo GeoRepository) *GeoService {
	return &GeoService{repo: repo}
}

func (s *GeoService) AddSearchHistory(search string) (int, error) {
	start := time.Now()

	secrhID, err := s.repo.AddSearchHistory(search)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("AddSearchHistory").Observe(duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return secrhID, nil
}

func (s *GeoService) AddAddress(data entities.Address) (int, error) {
	start := time.Now()

	addressID, err := s.repo.AddAddress(data)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("AddAddress").Observe(duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return addressID, nil
}

func (s *GeoService) AddHistorySearchAddress(id_search, id_address int) (int, error) {
	start := time.Now()

	id, err := s.repo.AddHistorySearchAddress(id_search, id_address)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("AddHistorySearchAddress").Observe(duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return id, nil
}

func (s *GeoService) GetAddressesBySearchData(search string) ([]entities.Address, error) {
	start := time.Now()

	arrAddress, err := s.repo.GetAddressesBySearchData(search)

	duration := time.Since(start).Seconds()
	metrics.DbDuration.WithLabelValues("GetAddressesBySearchData").Observe(duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения адресов: %v", err)
	}
	return arrAddress, nil
}
