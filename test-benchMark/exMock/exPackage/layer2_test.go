package exPackage_test

import (
	"golang-public/test-benchMark/exMock/exPackage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"1が返却される",
			func(t *testing.T) {
				e := exPackage.NewLayer2()

				assert.Equal(t, e.Get(), 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"trueが返却される",
			func(t *testing.T) {
				e := exPackage.NewLayer2()

				assert.Equal(t, e.Put(2), true)
			},
		},
		{
			"falseが返却される",
			func(t *testing.T) {
				e := exPackage.NewLayer2()

				assert.Equal(t, e.Put(3), false)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"trueが返却される",
			func(t *testing.T) {
				e := exPackage.NewLayer2()

				assert.Equal(t, e.Delete(3), true)
			},
		},
		{
			"falseが返却される",
			func(t *testing.T) {
				e := exPackage.NewLayer2()

				assert.Equal(t, e.Delete(4), false)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
