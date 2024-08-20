package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type sqlLite3 struct {
	db *sql.DB
}

func (s *sqlLite3) CreateUserTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    name TEXT,
    age INTEGER
)`)
	if err != nil {
		return err
	}
	log.Println("Создана таблица")
	return nil
}
func (s *sqlLite3) InsertUser(user User) error {
	_, err := s.db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}
	log.Printf("Добавлен пользователь с id=%d", user.ID)
	return nil
}
func (s *sqlLite3) SelectUser(id int) (User, error) {
	var user User

	rows, err := s.db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return User{}, err
		}
	} else {

		return User{}, fmt.Errorf("не найден пользователь с id=%d", id)
	}

	return user, nil
}
func (s *sqlLite3) UpdateUser(user User) error {
	_, err := s.db.Exec("UPDATE user SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}
	log.Printf("Обновление пользователя с id=%d", user.ID)
	return nil
}
func (s *sqlLite3) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return err
	}
	log.Printf("Удален пользователь с id=%d", id)
	return nil
}

func main() {

	database, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer database.Close()

	db := sqlLite3{database}

	err = db.CreateUserTable()
	if err != nil {
		fmt.Println(err)
	}

	userOne := User{
		ID:   1,
		Name: "Igor",
		Age:  29,
	}
	userTwo := User{
		ID:   2,
		Name: "Ivan",
		Age:  42,
	}
	userThree := User{
		ID:   3,
		Name: "Sergey",
		Age:  30,
	}

	err = db.InsertUser(userOne)
	err = db.InsertUser(userTwo)
	err = db.InsertUser(userThree)

	us, err := db.SelectUser(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID=1:", us)
	us, err = db.SelectUser(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID=2:", us)
	us, err = db.SelectUser(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID=3:", us)

	err = db.UpdateUser(User{ID: 2, Name: "Jax", Age: 87})
	us, err = db.SelectUser(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID=2 Update:", us)

	err = db.DeleteUser(1)
	fmt.Println("Delete User ID=1")

	us, err = db.SelectUser(1)
	if err != nil {
		fmt.Println(err)
	}

}
