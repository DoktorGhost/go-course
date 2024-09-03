package main

import (
	"testing"
)

// Функция для тестирования хэш-карты
func TestHashMap(t *testing.T) {
	tests := []struct {
		name     string
		hashFunc func(*HashMap)
	}{
		{"CRC32", WithHashCRC32()},
		{"CRC64", WithHashCRC64()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewHashMap(16, tt.hashFunc)

			// Тестирование установки и получения значения
			m.Set("key1", "value1")
			if value, ok := m.Get("key1"); !ok || value != "value1" {
				t.Errorf("expected value1, got %v", value)
			}

			// Тестирование получения несуществующего значения
			if _, ok := m.Get("nonexistent"); ok {
				t.Errorf("expected no value for nonexistent key")
			}
		})
	}
}

// Функция для тестирования функции hash-функций
func TestHashFunctions(t *testing.T) {
	tests := []struct {
		name     string
		hashFunc func(*HashMap)
	}{
		{"CRC32", WithHashCRC32()},
		{"CRC64", WithHashCRC64()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewHashMap(16, tt.hashFunc)
			m.Set("key", "value")
			if value, ok := m.Get("key"); !ok || value != "value" {
				t.Errorf("expected value, got %v", value)
			}
		})
	}
}

func BenchmarkHashMapCRC32(b *testing.B) {
	m := NewHashMap(16, WithHashCRC32())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		m.Get("key")
	}
}

func BenchmarkHashMapCRC64(b *testing.B) {
	m := NewHashMap(16, WithHashCRC64())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		m.Get("key")
	}
}
