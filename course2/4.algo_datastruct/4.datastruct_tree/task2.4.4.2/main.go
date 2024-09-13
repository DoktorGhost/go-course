package main

import (
	"fmt"
	"github.com/google/btree"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) Less(than btree.Item) bool {
	return u.ID < than.(User).ID
}

type BTree struct {
	tree *btree.BTree
}

func NewBTree(degree int) *BTree {
	return &BTree{
		tree: btree.New(degree),
	}
}

func (bt *BTree) Insert(user User) {
	bt.tree.ReplaceOrInsert(user)
}

func (bt *BTree) Search(id int) *User {

	item := bt.tree.Get(User{ID: id})
	if item != nil {
		user := item.(User)
		return &user
	}
	return nil
}

func main() {
	bt := NewBTree(2)
	users := []User{
		{1, "Alice", 30},
		{2, "Bob", 25},
		{3, "Charlie", 35},
	}
	for _, user := range users {
		bt.Insert(user)
	}
	if user := bt.Search(2); user != nil {
		fmt.Printf("Найден пользователь: %v\n", *user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
