package main

import (
	"testing"
)

// Benchmark functions

func BenchmarkEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GenerateUSER(100)
		EasyJSON(users)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GenerateUSER(100)
		JSON(users)
	}
}

func BenchmarkJSONiter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GenerateUSER(100)
		JSONiter(users)
	}
}
