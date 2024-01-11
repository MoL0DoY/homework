package main

import (
	"fmt"
)

const (
	COMPLEX = "complex"
	NUMERIC = "numeric"
)

type Calculator struct {
	Type     string
	Operator string
	A        IExpression
	B        IExpression
}

func (c *Calculator) Execute() (*IExpression, error) {
	switch c.Operator {
	case "+":
		return c.A.Add(c.B), nil
	case "-":
		return c.A.Sub(c.B), nil
	case "*":
		return c.A.Mul(c.B), nil
	case "/":
		return c.A.Div(c.B), nil
	}

	return nil, fmt.Errorf("ошибка оператора")
}

func complexCalculator() {

}
