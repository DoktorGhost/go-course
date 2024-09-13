package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

type User struct {
	ID   int
	Name string
	Age  int
}

func (t *BinaryTree) insert(user *User) *BinaryTree {
	if t.root == nil {
		t.root = &Node{
			index: user.ID,
			data:  user,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.insert(user)
	}
	return t
}

func (n *Node) insert(user *User) {
	if user.ID < n.index {
		if n.left == nil {
			n.left = &Node{
				index: user.ID,
				data:  user,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.insert(user)
		}
	} else {
		if n.right == nil {
			n.right = &Node{
				index: user.ID,
				data:  user,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.insert(user)
		}
	}
}

func (t *BinaryTree) search(key int) *User {
	if t.root == nil {
		return nil
	}
	return t.root.search(key)
}

func (n *Node) search(key int) *User {
	if n == nil {
		return nil
	}
	if n.index == key {
		return n.data
	}
	if key < n.index {
		return n.left.search(key)
	}

	return n.right.search(key)

}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func main() {
	bt := generateData(50)
	user := bt.search(30)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}

	user = bt.search(10)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}

/*
-Определены структуры данных User, Node и BinaryTree
-Реализован метод insert для вставки пользователей в бинарное дерево
-Реализован метод search для поиска пользователей по их идентификатору.
-Функция generateData генерирует случайные данные пользователей и вставляет их в бинарное дерево.
-Выводится информация о найденном пользователе, если он существует.
-Выводится сообщение «Пользователь не найден», если пользователь не найден.
*/
