package usecase

import (
	"geo/internal/entities"
	"geo/internal/service/cache"
	"geo/internal/service/dbservice"
	"geo/internal/service/geoservice"
	"log"
)

type UseCase struct {
	DBService    *dbservice.DBService
	CacheService *cache.CacheService
	GeoService   *geoservice.GeoService
}

func NewGeoUseCase(dbService *dbservice.DBService, cacheService *cache.CacheService, geoService *geoservice.GeoService) *UseCase {
	return &UseCase{DBService: dbService, CacheService: cacheService, GeoService: geoService}
}

// запрос данных кэш->бд->апи
func (uc *UseCase) SearchAddress(search string) ([]entities.Address, error) {
	var addresses []entities.Address
	//идем в кэш
	addresses, err := uc.CacheService.GetData(search)
	//если данные есть в кэше - возвращаем их
	if err == nil && len(addresses) > 0 {
		log.Println("Данные из кэша")
		return addresses, nil
	}
	//во всех других случаях идем в БД
	addresses, err = uc.DBService.GetAddressesBySearchData(search)
	//если из базы вернулся не пустой результат - то возвращаем его
	if err == nil && len(addresses) > 0 {
		log.Println("Данные из БД")
		//записываем сразу в кэш
		err = uc.CacheService.SaveData(search, addresses)
		if err != nil {
			log.Println("ошибка записи в кэш", err)
		}
		return addresses, nil
	}
	//в противном случае идем в АПИ
	resp, err := uc.GeoService.AddressSearch(entities.SearchRequest{Query: search})
	if err != nil {
		return nil, err
	}
	//записываем в кэш
	addresses = resp.Addresses
	err = uc.CacheService.SaveData(search, addresses)
	if err != nil {
		log.Println("ошибка записи в кэш", err)
	}
	//записываем в бд
	idSearch, err := uc.DBService.AddSearchHistory(search)
	if err != nil {
		log.Println("ошибка записи в бд", err)
	}
	for _, address := range resp.Addresses {
		idAddress, err := uc.DBService.AddAddress(address)
		if err != nil {
			log.Println("ошибка записи в бд", err)
		}
		_, err = uc.DBService.AddHistorySearchAddress(idSearch, idAddress)
		if err != nil {
			log.Println("ошибка записи в бд", err)
		}
	}

	return addresses, nil
}

func (uc *UseCase) PerformGeocode(input entities.GeocodeRequest) (entities.Response, error) {
	return uc.GeoService.GeoCode(input)
}
