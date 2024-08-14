package main

import "fmt"

func main() {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(sum(xs))     //36
	fmt.Println(average(xs)) //4.5

	ys := [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}

	fmt.Println(averageFloat(ys)) //5
	fmt.Println(reverse(xs))      // 8 7 6 5 4 3 2 1
}

func sum(xs [8]int) int {
	var sum int
	for _, x := range xs {
		sum += x
	}
	return sum
}

func average(xs [8]int) float64 {
	var sum int
	for _, x := range xs {
		sum += x
	}
	sumF := float64(sum)
	lenF := float64(len(xs))

	return sumF / lenF
}

func averageFloat(ys [8]float64) float64 {
	var sum float64
	for _, x := range ys {
		sum += x
	}

	lenF := float64(len(ys))

	return sum / lenF
}

func reverse(xs [8]int) [8]int {
	for i := 0; i < len(xs)/2; i++ {
		j := len(xs) - i - 1
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}
