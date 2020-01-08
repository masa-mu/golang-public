package exPackage

type exStruct1 struct {
	Layer2 Layer2Interface
}

type Layer1Interface interface {
	Receiver() bool
}

func NewLayer1(Layer2 Layer2Interface) Layer1Interface {
	return &exStruct1{Layer2}
}

func (e *exStruct1) Receiver() bool {
	if e.Layer2.Get() != 1 {
		return false
	}

	if e.Layer2.Put(2) == false {
		return false
	}

	if e.Layer2.Delete(3) == false {
		return false
	}

	return true
}
