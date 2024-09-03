package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type SQLiteGenerator struct{}

func (sg SQLiteGenerator) CreateTableSQL(model Tabler) string {
	modelType := reflect.TypeOf(model).Elem()
	if modelType.Kind() != reflect.Struct {
		return ""
	}

	tableName := model.TableName()
	sql := fmt.Sprintf("CREATE TABLE %s (\n", tableName)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")

		sql += fmt.Sprintf("%s %s,\n", dbField, dbType)
	}
	sql = sql[:len(sql)-2]
	sql += ");\n"
	return sql
}

func (sg SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	if modelType.Kind() != reflect.Struct {
		return ""
	}

	tableName := model.TableName()
	sql := fmt.Sprintf("INSERT INTO %s (", tableName)

	var fields []string
	var values []string

	for i := 1; i < modelType.NumField(); i++ { // Пропускаем поле ID
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")

		fields = append(fields, dbField)

		fieldValue := modelValue.Field(i).String()
		valueStr := fmt.Sprintf("'%s'", fieldValue)

		values = append(values, valueStr)
	}

	sql += fmt.Sprintf("%s) VALUES (%s);", strings.Join(fields, ", "), strings.Join(values, ", "))
	return sql
}

type GoFakeitGenerator struct{}

func (fg GoFakeitGenerator) GenerateFakeUser() User {
	u := User{}
	u.FirstName = gofakeit.FirstName()
	u.LastName = gofakeit.LastName()
	u.Email = gofakeit.Email()
	return u
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)
	fakeUser := fakeDataGenerator.GenerateFakeUser()
	query := sqlGenerator.CreateInsertSQL(&fakeUser)
	fmt.Println(query)
}
