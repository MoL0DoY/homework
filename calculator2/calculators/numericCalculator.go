package calculators

import (
	"fmt"
	"go2/calculator2/models"
)

type NumericCalculator struct {
	Type     string
	Operator string
	A        models.RealNumber
	B        models.RealNumber
}

func (c *NumericCalculator) Execute(operator string) (string, error) {
	var result models.RealNumber
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
