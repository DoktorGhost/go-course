package psg

import (
	"context"
	"fmt"
	"geoservice/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: db}
}

func (s *UsersRepository) CreateUser(user entities.UserType) (int, error) {
	var id int
	query := `INSERT INTO users.users (username, password) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, user.Username, user.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *UsersRepository) GetUser(username string) (entities.UserType, error) {
	var result entities.UserType
	query := `SELECT id, username, password FROM users.users WHERE username = $1`
	err := s.db.QueryRow(context.Background(), query, username).Scan(&result.ID, &result.Username, &result.Password)

	if err != nil {
		return entities.UserType{}, fmt.Errorf("[PostgreSQL] Error: %v", err)
	}

	return result, nil
}
