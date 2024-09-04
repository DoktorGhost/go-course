package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"reflect"
	"strings"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

func (sg SQLiteGenerator) CreateTableSQL(table Tabler) string {
	modelType := reflect.TypeOf(table).Elem()
	if modelType.Kind() != reflect.Struct {
		return ""
	}

	tableName := table.TableName()
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", tableName)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")

		sql += fmt.Sprintf("%s %s,\n", dbField, dbType)
	}
	sql = strings.TrimSuffix(sql, ",\n") + ");\n"
	return sql
}

func (sg SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	modelType := reflect.TypeOf(model).Elem()
	if modelType.Kind() != reflect.Struct {
		return ""
	}

	tableName := model.TableName()
	sql := fmt.Sprintf("INSERT INTO %s (", tableName)

	var fields []string
	var placeholders []string

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")

		if dbField == "" {
			continue
		}

		fields = append(fields, dbField)
		placeholders = append(placeholders, "?")
	}

	sql += strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ");\n"
	return sql
}

type Tabler interface {
	TableName() string
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m *Migrator) Migrate(tablers ...Tabler) error {
	for _, tabler := range tablers {
		// Создание таблицы
		createTableSQL := m.sqlGenerator.CreateTableSQL(tabler)
		if _, err := m.db.Exec(createTableSQL); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
		// Вставка данных
		insertSQL := m.sqlGenerator.CreateInsertSQL(tabler)
		if _, err := m.db.Exec(insertSQL, 1, "John", "Doe", "john.doe@example.com"); err != nil {
			return fmt.Errorf("failed to insert data: %w", err)
		}
	}
	return nil
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	sqlGenerator := &SQLiteGenerator{}
	migrator := NewMigrator(db, sqlGenerator)
	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
