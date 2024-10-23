package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"user/internal/entities"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: db}
}

func (s *UsersRepository) CreateUser(user entities.User) (int, error) {
	var id int
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, user.Email, user.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *UsersRepository) GetUser(email string) (entities.User, error) {
	var result entities.User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := s.db.QueryRow(context.Background(), query, email).Scan(&result.ID, &result.Email, &result.Password)

	if err != nil {
		return entities.User{}, fmt.Errorf("[PostgreSQL] Error: %v", err)
	}

	return result, nil
}

func (s *UsersRepository) GetAllUser() ([]entities.User, error) {
	var users []entities.User
	query := `SELECT id, email, password FROM users`

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("[PostgreSQL] Error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("[PostgreSQL] Error scanning user: %v", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[PostgreSQL] Error during rows iteration: %v", err)
	}

	return users, nil
}
