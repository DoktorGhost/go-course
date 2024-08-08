package main

import "math/rand"

func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func generateTestData(len int) []float64 {
	testData := make([]float64, len)
	for i := 0; i < len; i++ {
		testData[i] = float64(rand.Intn(100))
	}
	return testData
}
