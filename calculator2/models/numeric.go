package models

type RealNumber struct {
	value float64
}

func (n *RealNumber) Add(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value + a.Value()}
}

func (n *RealNumber) Sub(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value - a.Value()}
}

func (n *RealNumber) Mul(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value - a.Value()}
}

func (n *RealNumber) Div(a RealNumber) *RealNumber {
	if a.Value() == 0 {
		panic(a)
	}
	return &RealNumber{value: n.value / a.Value()}
}

func (n *RealNumber) Value() float64 {
	return n.value
}

func (n *RealNumber) SetValue(value float64) {
	n.value = value
}
