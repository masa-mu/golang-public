package exPackage

type exStruct2 struct {
}

type Layer2Interface interface {
	Get() int
	Put(param int) bool
	Delete(key int) bool
}

func NewLayer2() Layer2Interface {
	return &exStruct2{}
}

func (e *exStruct2) Get() int {
	return 1
}

func (e *exStruct2) Put(param int) bool {
	if param == 2 {
		return true
	} else {
		return false
	}
}

func (e *exStruct2) Delete(key int) bool {
	if key == 3 {
		return true
	} else {
		return false
	}
}
