package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Node struct {
	data *Commit
	prev *Node
	next *Node
}
type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type Commit struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Date    string `json:"date"`
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

// LoadData loads data from a JSON file at the given path into the list.
func (d *DoubleLinkedList) LoadData(path string) error {
	if path == "" {
		return errors.New("nil путь файла")
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit
	if err := json.Unmarshal(bytes, &commits); err != nil {
		return err
	}
	QuickSort(commits)
	d.Init(commits)
	return nil
}

// Инициализирует двусвязный список массивом объектов Commit
func (d *DoubleLinkedList) Init(c []Commit) {
	if len(c) < 1 {
		d.head = nil
		d.tail = nil
		d.curr = nil
		d.len = 0
	}
	for _, commit := range c {
		node := &Node{data: &commit, prev: nil, next: nil}

		if d.head == nil {
			d.head = node
			d.len = 1
			d.curr = node
			d.tail = node
		} else {
			current := d.head
			for current.next != nil {
				current = current.next
			}
			current.next = node
			node.prev = current
			d.len++
			d.curr = node
			d.tail = node
		}
	}
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n < 0 {
		return errors.New("n не должно быть < 0")
	}
	if n > d.len-1 {
		return errors.New("n больше длины списка")
	}

	if n == d.len-1 {
		d.curr = d.tail
		return nil
	}

	node := d.head
	for i := 1; i <= n; i++ {
		next := node.next
		node = next
	}
	d.curr = node
	return nil
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	if d.curr.next == nil {
		return nil
	}
	d.curr = d.curr.next
	return d.curr
}

// Next получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {

	d.curr = d.curr.prev
	return d.curr
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}
	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	d.len++
	return nil
}

func (d *DoubleLinkedList) Push(c Commit) error {
	if d.tail == nil {
		return errors.New("нет последнего элемента")
	}
	node := &Node{data: &c, prev: d.tail, next: nil}
	d.tail = node
	d.len++

	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n < 0 {
		return errors.New("n < 0")
	}
	if n >= d.len {
		return errors.New("n больше длины списка")
	}
	if n == d.len-1 {
		d.tail = d.tail.prev
		d.len--
		return nil
	}
	if n == 0 {
		d.head = d.head.next
		d.len--
		return nil
	}

	node := d.head
	for i := 0; i < n; i++ {
		next := node.next
		node = next
	}
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
	d.len--

	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return errors.New("нет текущего элемента")
	}
	if d.curr == d.head {
		d.head = d.head.prev
		d.len--
		return nil
	}
	if d.curr == d.tail {
		d.tail = d.tail.prev
		d.len--
		return nil
	}

	prev := d.curr.prev
	next := d.curr.next

	prev.next = next
	next.prev = prev
	d.len--
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	if d.curr == nil {
		return 0, errors.New("нет текущего элемента")
	}
	if d.curr == d.head {
		return 1, nil
	}
	if d.curr == d.tail {
		return d.len, nil
	}

	node := d.head
	idx := 1

	for i := 0; i < d.len; i++ {
		idx++
		if node.next == d.curr {
			return idx, nil
		}
	}

	return 0, errors.New("не найден индекс текущего элемента")
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n > d.len {
		return nil, errors.New("индекс выходит за границы списка")
	}
	if n == 1 {
		return d.head, nil
	}
	if n == d.len {
		return d.head, nil
	}
	node := d.head
	idx := 1

	for i := 1; i < n; i++ {
		idx++
		node = node.next
	}
	return node, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	tail := d.tail
	newTail := d.tail.prev
	newTail.next = nil
	d.tail = newTail
	d.len--

	return tail
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	head := d.head
	newHead := d.head.next
	newHead.prev = nil
	d.head = newHead
	d.len--
	return head
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	node := d.head
	for i := 0; i < d.len; i++ {
		if node.data.UUID == uuID {
			return node
		}
		node = node.next
	}
	return nil
}

func (d *DoubleLinkedList) Search(message string) *Node {
	node := d.head
	for i := 0; i < d.len; i++ {
		if node.data.Message == message {
			return node
		}
		node = node.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	head := d.head
	tail := d.tail

	reverseDoubleLinkedList := &DoubleLinkedList{}

	reverseDoubleLinkedList.head = tail
	reverseDoubleLinkedList.tail = head

	node := reverseDoubleLinkedList.head

	for i := 0; i < d.len; i++ {
		node.next, node.prev = node.prev, node.next
		node = node.next
	}

	return reverseDoubleLinkedList
}

// сортировка
func quickSortHelper(commits []Commit, low, high int) {
	if low < high {
		p := partition(commits, low, high)
		quickSortHelper(commits, low, p-1)
		quickSortHelper(commits, p+1, high)
	}
}

func partition(commits []Commit, low, high int) int {
	pivot := commits[high]
	datePivot, _ := time.Parse("2006-01-02", pivot.Date)

	i := low - 1

	for j := low; j < high; j++ {
		date, _ := time.Parse("2006-01-02", commits[j].Date)
		if date.Before(datePivot) {
			i++
			commits[i], commits[j] = commits[j], commits[i]
		}
	}

	commits[i+1], commits[high] = commits[high], commits[i+1]
	return i + 1
}

func QuickSort(commits []Commit) {
	quickSortHelper(commits, 0, len(commits)-1)
}
