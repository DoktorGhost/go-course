package main

import (
	"testing"
)

func TestUser_TableName_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TableName() did not panic")
		}
	}()

	user := User{}
	user.TableName()
}

func TestUser_CreateTableSQL_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateTableSQL() did not panic")
		}
	}()

	user := User{}
	user.CreateTableSQL(user)
}

func TestUser_GenerateFakeUser_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GenerateFakeUser() did not panic")
		}
	}()

	user := User{}
	user.GenerateFakeUser()
}
