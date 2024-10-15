package geo_usecase

import (
	"errors"
	"fmt"
	"geoservice/internal/entities"
	"geoservice/internal/services/geo_services"
	"log"
)

type GeoUseCase struct {
	geoService *geo_services.GeoService
}

func NewGeoUseCase(geoService *geo_services.GeoService) *GeoUseCase {
	return &GeoUseCase{geoService: geoService}
}

func (uc *GeoUseCase) AddSearch(search string, arrAddress []entities.Address) error {
	//проверка, что такого запроса не было
	arr, err := uc.geoService.GetAddressesBySearchData(search)
	if err != nil {
		return errors.New("ошибка GetAddressesBySearchData: " + err.Error())
	}
	//результат есть, записывать в БД не надо
	if len(arr) != 0 {
		log.Println("GetAddressesBySearchData вернул не пустой результат")
		return nil
	}

	//записываем все в БД
	idSearch, err := uc.geoService.AddSearchHistory(search)
	if err != nil {
		return fmt.Errorf("ошибка записи AddSearchHistory: " + err.Error())
	}

	for _, value := range arrAddress {
		idAddress, err := uc.geoService.AddAddress(value)
		if err != nil {
			return fmt.Errorf("ошибка записи AddAddress: " + err.Error())
		}
		_, err = uc.geoService.AddHistorySearchAddress(idSearch, idAddress)
		if err != nil {
			return fmt.Errorf("ошибка записи AddHistorySearchAddress: " + err.Error())
		}
	}

	return nil
}

func (uc *GeoUseCase) GetSearch(search string) ([]entities.Address, error) {
	arr, err := uc.geoService.GetAddressesBySearchData(search)
	if err != nil {
		return nil, errors.New("ошибка GetAddressesBySearchData: " + err.Error())
	}

	return arr, nil
}
