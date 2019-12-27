package exTest

// 適当にスライスを作成(structでも何でも良い）
type swrap []int

// swrapの持つレシーバを定義
type SWrap interface {
	Add(a int)
	Len() int
	Get() []int
}

// 外部パッケージから呼び出すメソッド
func NewSample() SWrap {
	return &swrap{}
}

// 適当なレシーバ1
func (sw *swrap) Add(a int) {
	*sw = append(*sw, a)
}

// 適当なレシーバ2
func (sw *swrap) Len() int {
	return len(*sw)
}

// 適当なレシーバ3
func (sw *swrap) Get() []int {
	return *sw
}
