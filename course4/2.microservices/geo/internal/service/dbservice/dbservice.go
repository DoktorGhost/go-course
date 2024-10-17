package dbservice

import (
	"fmt"
	"geo/internal/entities"
)

type DBRepository interface {
	AddSearchHistory(search string) (int, error)
	AddAddress(data entities.Address) (int, error)
	AddHistorySearchAddress(id_search, id_address int) (int, error)
	GetAddressesBySearchData(search string) ([]entities.Address, error)
}

type DBService struct {
	repo DBRepository
}

func NewDBService(repo DBRepository) *DBService {
	return &DBService{repo: repo}
}

func (s *DBService) AddSearchHistory(search string) (int, error) {
	secrhID, err := s.repo.AddSearchHistory(search)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return secrhID, nil
}

func (s *DBService) AddAddress(data entities.Address) (int, error) {
	addressID, err := s.repo.AddAddress(data)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return addressID, nil
}

func (s *DBService) AddHistorySearchAddress(id_search, id_address int) (int, error) {
	id, err := s.repo.AddHistorySearchAddress(id_search, id_address)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return id, nil
}

func (s *DBService) GetAddressesBySearchData(search string) ([]entities.Address, error) {
	arrAddress, err := s.repo.GetAddressesBySearchData(search)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения адресов: %v", err)
	}
	return arrAddress, nil
}
