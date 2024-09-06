package main

import (
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	//benchmark code
	for i := 0; i < b.N; i++ {
		p = &Person{}
		p.Age = i
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	//benchmark code
	for i := 0; i < b.N; i++ {
		p = personPool.Get().(*Person)
		p.Age = i
		personPool.Put(p)
	}
}
