package main

import "fmt"

type Calculator struct {
}

func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на ноль недопустимо")
	}
	return a / b, nil
}

type ComplexCalculator struct{}

func (c ComplexCalculator) Add(a, b complex128) complex128 {
	return a + b
}

func (c ComplexCalculator) Subtract(a, b complex128) complex128 {
	return a - b
}

func (c ComplexCalculator) Multiply(a, b complex128) complex128 {
	return a * b
}

func (c ComplexCalculator) Divide(a, b complex128) (complex128, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на ноль недопустимо")
	}
	return a / b, nil

}
