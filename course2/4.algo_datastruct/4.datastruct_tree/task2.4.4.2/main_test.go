package main

import (
	"reflect"
	"testing"
)

func TestBTree_InsertAndSearch(t *testing.T) {
	bt := NewBTree(2)

	users := []User{
		{1, "Alice", 30},
		{2, "Bob", 25},
		{3, "Charlie", 35},
	}

	for _, user := range users {
		bt.Insert(user)
	}

	tests := []struct {
		id       int
		expected *User
	}{
		{id: 1, expected: &User{ID: 1, Name: "Alice", Age: 30}},
		{id: 2, expected: &User{ID: 2, Name: "Bob", Age: 25}},
		{id: 3, expected: &User{ID: 3, Name: "Charlie", Age: 35}},
	}

	for _, tt := range tests {
		t.Run(tt.expected.Name, func(t *testing.T) {
			result := bt.Search(tt.id)

			if result == nil && tt.expected == nil {
				return
			}
			if result == nil || !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Search(%d) = %v, ожидалось %v", tt.id, result, tt.expected)
			}
		})
	}
}

func TestBTree_Insert_Duplicate(t *testing.T) {
	bt := NewBTree(2)

	// Вставка пользователя с одним и тем же ID дважды
	user := User{ID: 1, Name: "Alice", Age: 30}
	bt.Insert(user)
	bt.Insert(User{ID: 1, Name: "Alice", Age: 31}) // Вставляем пользователя с таким же ID, но другим Age

	// Ожидаем, что Age должен обновиться
	expected := &User{ID: 1, Name: "Alice", Age: 31}
	result := bt.Search(1)
	if result == nil || !reflect.DeepEqual(result, expected) {
		t.Errorf("После повторной вставки Search(1) = %v, ожидалось %v", result, expected)
	}
}
