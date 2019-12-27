package exTest_test

import (
	"golang-public/test-benchMark/exTest"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"1~10のランダムな数字が3個追加される",
			func(t *testing.T) {
				Sw := exTest.NewSample()
				for i := 0; i < 3; i++ {
					Sw.Add(rand.Intn(10))
				}
				assert.Equal(t, len(Sw.Get()), 3)
			},
		},
		{
			"1~10のランダムな数字が5個追加される",
			func(t *testing.T) {
				Sw := exTest.NewSample()
				for i := 0; i < 5; i++ {
					Sw.Add(rand.Intn(10))
				}
				assert.Equal(t, len(Sw.Get()), 5)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"1~10のランダムな数字が3個追加される",
			func(t *testing.T) {
				Sw := exTest.NewSample()
				for i := 0; i < 3; i++ {
					Sw.Add(rand.Intn(10))
				}
				assert.Equal(t, Sw.Len(), 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
