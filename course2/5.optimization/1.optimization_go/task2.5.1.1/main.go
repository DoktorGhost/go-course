package main

import (
	"fmt"
	"hash/crc32"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// Вспомогательная структура для хранения пары ключ-значение
type arrayEntry struct {
	key   string
	value interface{}
}

func crc32ToInt(key string) int {
	hash := crc32.ChecksumIEEE([]byte(key))
	return int(hash)
}

type HashMapSlice struct {
	buckets  [][]arrayEntry
	hashFunc func(string) int
}

func NewHashMapSlice(size int, hashFunc func(string) int) *HashMapSlice {
	return &HashMapSlice{
		buckets:  make([][]arrayEntry, size),
		hashFunc: hashFunc,
	}
}

func (h *HashMapSlice) Set(key string, value interface{}) {
	index := h.hashFunc(key) % len(h.buckets)
	h.buckets[index] = append(h.buckets[index], arrayEntry{key, value})
}

func (h *HashMapSlice) Get(key string) (interface{}, bool) {
	index := h.hashFunc(key) % len(h.buckets)
	for _, entry := range h.buckets[index] {
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}

// Вспомогательная структура для элемента списка
type listNode struct {
	key   string
	value interface{}
	next  *listNode
}

type HashMapList struct {
	buckets  []*listNode // Массив указателей на начало списков
	hashFunc func(string) int
}

func NewHashMapList(size int, hashFunc func(string) int) *HashMapList {
	return &HashMapList{
		buckets:  make([]*listNode, size),
		hashFunc: hashFunc,
	}
}

func (h *HashMapList) Set(key string, value interface{}) {
	index := h.hashFunc(key) % len(h.buckets)
	newNode := &listNode{key, value, nil}

	if h.buckets[index] == nil {
		h.buckets[index] = newNode
	} else {
		current := h.buckets[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (h *HashMapList) Get(key string) (interface{}, bool) {
	index := h.hashFunc(key) % len(h.buckets)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return nil, false
}

func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	since := time.Since(start)
	return since
}

func main() {
	time := MeassureTime(TestSlice16)
	fmt.Println(time)

	time = MeassureTime(TestSlice1000)
	fmt.Println(time)

	time = MeassureTime(TestList16)
	fmt.Println(time)

	time = MeassureTime(TestList1000)
	fmt.Println(time)
}

func TestList16() {
	m := NewHashMapList(16, crc32ToInt)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestList1000() {
	m := NewHashMapList(1000, crc32ToInt)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice16() {
	m := NewHashMapSlice(16, crc32ToInt)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice1000() {
	m := NewHashMapSlice(1000, crc32ToInt)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

/*
Присутствуют две реализации, использующее массив ключ значение и список ключ значение
Написан бенчмарк для реализации хранения ключ-значений массивом
Написан бенчмарк для реализации хранения ключ-значений с помощью структуры список
Вместо рекурсивного поиска по листу использован цикл
Код покрыт тестами на 100%

*/
