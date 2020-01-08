package exMock

import "golang-public/test-benchMark/exMock/exPackage"

type Layer2 struct {
	exPackage.Layer2Interface

	FakeGet    func() int
	FakePut    func(param int) bool
	FakeDelete func(key int) bool
}

func (l Layer2) Get() int {
	return l.FakeGet()
}

func (l Layer2) Put(param int) bool {
	return l.FakePut(param)
}

func (l Layer2) Delete(key int) bool {
	return l.FakeDelete(key)
}
