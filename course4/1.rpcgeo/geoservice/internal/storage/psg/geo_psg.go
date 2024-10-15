package psg

import (
	"context"
	"fmt"
	"geoservice/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GeoRepository struct {
	db *pgxpool.Pool
}

func NewGeoRepository(db *pgxpool.Pool) *GeoRepository {
	return &GeoRepository{db: db}
}

func (s *GeoRepository) AddSearchHistory(search string) (int, error) {
	var id int
	query := `INSERT INTO geo.search_history (data) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, search).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *GeoRepository) AddAddress(data entities.Address) (int, error) {
	var id int
	query := `INSERT INTO geo.address (city, street, house, lat, lon) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, data.City, data.Street, data.House, data.Lat, data.Lon).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *GeoRepository) AddHistorySearchAddress(id_search, id_address int) (int, error) {
	var id int
	query := `INSERT INTO geo.history_search_address (id_search, id_address) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, id_search, id_address).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *GeoRepository) GetAddressesBySearchData(search string) ([]entities.Address, error) {
	var addresses []entities.Address

	// Запрос для получения адресов по данным поиска
	query := `
		SELECT a.id, a.city, a.street, a.house, a.lat, a.lon
		FROM geo.address a
		JOIN geo.history_search_address hsa ON a.id = hsa.id_address
		WHERE hsa.id_search IN (
					SELECT id
					FROM geo.search_history
					WHERE levenshtein(data, $1) <= LENGTH($1) * 0.3
		);
	`

	// Выполняем запрос
	rows, err := s.db.Query(context.Background(), query, search)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close() // Закрываем rows после завершения работы с ними

	// Читаем результаты
	for rows.Next() {
		var address entities.Address
		if err := rows.Scan(&address.ID, &address.City, &address.Street, &address.House, &address.Lat, &address.Lon); err != nil {
			return nil, fmt.Errorf("ошибка сканирования результата: %v", err)
		}
		addresses = append(addresses, address)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обходе результатов: %v", err)
	}

	return addresses, nil // Возвращаем массив адресов
}
