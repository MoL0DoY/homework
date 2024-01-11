package calculators

import (
	"fmt"
	"go2/calculator2/models"
)

type ComplexCalculator struct {
	Type     string
	Operator string
	A        models.Complex
	B        models.Complex
}

// Execute implements Calculator.
func (c *ComplexCalculator) Execute(operator string) (string, error) {
	var result models.Complex
	switch operator {
	case "+":
		result = *c.A.Add(c.B)
	case "-":
		result = *c.A.Sub(c.B)
	case "*":
		result = *c.A.Mul(c.B)
	case "/":
		result = *c.A.Div(c.B)
	default:
		return "", fmt.Errorf("Некорректный оператор")
	}
	return result.ToString(), nil
}
