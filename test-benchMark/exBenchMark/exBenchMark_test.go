package exBenchMark_test

import (
	"golang-public/test-benchMark/exBenchMark"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	Sw := exBenchMark.NewSample()

	for i := 0; i < b.N; i++ {
		Sw.Add(b.N)
	}
}

func BenchmarkLen(b *testing.B) {
	Sw := exBenchMark.NewSample()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sw.Len()
	}
}
