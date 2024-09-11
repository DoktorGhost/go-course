package main

import "fmt"

func GeneralSort(arr []int) {
	if len(arr) <= 10 {
		insertionSort(arr)
	} else if len(arr) <= 100 {
		selectionSort(arr)
	} else {
		arr = mergeSort(arr)
		copy(arr, mergeSort(arr))
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)
	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quicksort(data)
	fmt.Println("Sorted by Quicksort: ", sortedData)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", data)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(quicksort(left), pivot), quicksort(right)...)
}
