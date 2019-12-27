package sample_test

import (
	"golang-public/bench-mark/sample"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	Sw := sample.NewSample()

	for i := 0; i < b.N; i++ {
		Sw.Add(b.N)
	}
}

func BenchmarkLen(b *testing.B) {
	Sw := sample.NewSample()
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		Sw.Len()
	}
}

func TestAdd(b *testing.T) {
	Sw := sample.NewSample()
	for i := 0; i < 3; i++ {
		Sw.Add(1)
	}
}

func TestLen(b *testing.T) {
	Sw := sample.NewSample()

	for i := 0; i < 200; i++ {
		Sw.Len()
	}
}
