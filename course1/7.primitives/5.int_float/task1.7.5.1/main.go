package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
	var number uint32
	//
	numberUint, err := strconv.ParseUint(binary, 2, 32)
	if err != nil {
		return 0
	}
	number = uint32(numberUint)
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}
