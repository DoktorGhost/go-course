package main

import (
	"fmt"
	"sort"
)

func main() {
	intArr := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
	floatArr := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}

	sortedIntDesc := sortDescInt(intArr)
	sortedIntAsc := sortAscInt(intArr)
	sortedFloatDesc := sortDescFloat(floatArr)
	sortedFloatAsc := sortAscFloat(floatArr)

	fmt.Println("Sorted Int Array (Descending):", sortedIntDesc)
	fmt.Println("Sorted Int Array (Ascending):", sortedIntAsc)
	fmt.Println("Sorted Float Array (Descending):", sortedFloatDesc)
	fmt.Println("Sorted Float Array (Ascending):", sortedFloatAsc)
}

func sortDescInt(intArr [8]int) [8]int {
	slice := intArr[:]
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return intArr
}

func sortAscInt(intArr [8]int) [8]int {
	slice := intArr[:]
	sort.Sort(sort.IntSlice(slice))
	return intArr
}
func sortDescFloat(floatArr [8]float64) [8]float64 {
	slice := floatArr[:]
	sort.Sort(sort.Reverse(sort.Float64Slice(slice)))
	return floatArr
}
func sortAscFloat(floatArr [8]float64) [8]float64 {
	slice := floatArr[:]
	sort.Sort(sort.Float64Slice(slice))
	return floatArr
}
