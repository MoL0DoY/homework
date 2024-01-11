package models

import (
	"fmt"
	"strconv"
)

type RealNumber struct {
	value float64
}

func (n *RealNumber) Add(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value + a.value}
}

func (n *RealNumber) Sub(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value - a.value}
}

func (n *RealNumber) Mul(a RealNumber) *RealNumber {
	return &RealNumber{value: n.value - a.value}
}

func (n *RealNumber) Div(a RealNumber) *RealNumber {
	if a.value == 0 {
		panic(a)
	}
	return &RealNumber{value: n.value / a.value}
}

func (n *RealNumber) SetValue(value float64) {
	n.value = value
}

func (n *RealNumber) ToString() string {
	return fmt.Sprintf("%s", strconv.FormatFloat(n.value, 'f', 2, 64))
}
