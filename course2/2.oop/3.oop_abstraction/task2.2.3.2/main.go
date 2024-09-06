package main

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string {
	//TODO implement me
	panic("implement me")
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(model Tabler) string
}

func (u User) CreateTableSQL(model Tabler) string {
	//TODO implement me
	panic("implement me")
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (u User) GenerateFakeUser() User {
	//TODO implement me
	panic("implement me")
}
