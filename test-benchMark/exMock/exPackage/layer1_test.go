package exPackage_test

import (
	"golang-public/test-benchMark/exMock/exMock"
	"golang-public/test-benchMark/exMock/exPackage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiver(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{
			"Getでfalseが返却される",
			func(t *testing.T) {
				layer2Mock := exMock.Layer2{
					FakeGet: func() int {
						return 2
					},
				}

				layer1 := exPackage.NewLayer1(layer2Mock)

				assert.Equal(t, layer1.Receiver(), false)
			},
		},
		{
			"Putでfalseが返却される",
			func(t *testing.T) {
				layer2Mock := exMock.Layer2{
					FakeGet: func() int {
						return 1
					},
					FakePut: func(param int) bool {
						return false
					},
				}

				layer1 := exPackage.NewLayer1(layer2Mock)

				assert.Equal(t, layer1.Receiver(), false)
			},
		},
		{
			"Deleteでfalseが返却される",
			func(t *testing.T) {
				layer2Mock := exMock.Layer2{
					FakeGet: func() int {
						return 1
					},
					FakePut: func(param int) bool {
						return true
					},
					FakeDelete: func(key int) bool {
						return false
					},
				}

				layer1 := exPackage.NewLayer1(layer2Mock)

				assert.Equal(t, layer1.Receiver(), false)
			},
		},
		{
			"すべてが成功してtrueが返却される",
			func(t *testing.T) {
				layer2Mock := exMock.Layer2{
					FakeGet: func() int {
						return 1
					},
					FakePut: func(param int) bool {
						return true
					},
					FakeDelete: func(key int) bool {
						return true
					},
				}

				layer1 := exPackage.NewLayer1(layer2Mock)

				assert.Equal(t, layer1.Receiver(), true)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
