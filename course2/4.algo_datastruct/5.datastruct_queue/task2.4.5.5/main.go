package main

import "fmt"

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type RingBuffer struct {
	buffer chan int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{make(chan int, size)}
}

func (rb *RingBuffer) Add(val int) {
	if len(rb.buffer) == cap(rb.buffer) {
		<-rb.buffer
	}
	rb.buffer <- val
}

func (rb *RingBuffer) Get() (int, bool) {
	if len(rb.buffer) == 0 {
		return 0, false
	}
	val := <-rb.buffer
	return val, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4) // Перезаписывает значение 1
	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val) // Выводит: 2, 3, 4
	}
	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty") // Выводит: Buffer is empty
	}
}
