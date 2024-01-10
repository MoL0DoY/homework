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

type СomplexNumber struct {
	value1 complex128
}

func (n *СomplexNumber) Add(a СomplexNumber) *СomplexNumber {
	return &СomplexNumber{value1: n.value1 + a.Value1()}
}

func (n *СomplexNumber) Sub(a СomplexNumber) *СomplexNumber {
	return &СomplexNumber{value1: n.value1 - a.Value1()}
}

func (n *СomplexNumber) Mul(a СomplexNumber) *СomplexNumber {
	return &СomplexNumber{value1: n.value1 - a.Value1()}
}

func (n *СomplexNumber) Div(a СomplexNumber) *СomplexNumber {
	if a.Value1() == 0 {
		panic(a)
	}
	return &СomplexNumber{value1: n.value1 / a.Value1()}
}

func (n *СomplexNumber) Value1() complex128 {
	return n.value1
}

func (n *СomplexNumber) SetValue(value1 complex128) {
	n.value1 = value1
}
