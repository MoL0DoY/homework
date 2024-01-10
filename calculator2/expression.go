package main

type IExpression interface {
	Add(a IExpression) *IExpression
	Sub(a IExpression) *IExpression
	Mul(a IExpression) *IExpression
	Div(a IExpression) *IExpression
	Value() float64
	Value1() complex128
}
