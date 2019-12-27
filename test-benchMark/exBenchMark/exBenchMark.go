package exBenchMark

type swrap []int

type SWrap interface {
	Add(a int)
	Len() int
}

func NewSample() SWrap {
	return &swrap{}
}

func (sw *swrap) Add(a int) {
	*sw = append(*sw, a)
}

func (sw *swrap) Len() int {
	return len(*sw)
}
