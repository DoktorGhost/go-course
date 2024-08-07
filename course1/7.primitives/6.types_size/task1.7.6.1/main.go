package main

import (
	"fmt"
	"unsafe"
)

func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfStruct(n struct{}) int {
	return int(unsafe.Sizeof(n))
}

func main() {
	var (
		a bool
		b int
		c int8
		d int16
		e int32
		f int64
		g uint
		h uint8
		s struct{}
	)

	fmt.Println("bool size: ", sizeOfBool(a))
	fmt.Println("int size: ", sizeOfInt(b))
	fmt.Println("int8 size: ", sizeOfInt8(c))
	fmt.Println("int16 size: ", sizeOfInt16(d))
	fmt.Println("int32 size: ", sizeOfInt32(e))
	fmt.Println("int64 size: ", sizeOfInt64(f))
	fmt.Println("uint size: ", sizeOfUint(g))
	fmt.Println("uint8 size: ", sizeOfUint8(h))
	fmt.Println("nil struct size: ", sizeOfStruct(s))
}
