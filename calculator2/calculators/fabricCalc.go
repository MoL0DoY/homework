package calculators

import (
	"fmt"
	"go2/calculator2/models"
)

const (
	COMPLEX = "complex"
	NUMERIC = "numeric"
)

func NewCalculator(typeCalc string) (Calculator, error) {
	if typeCalc == NUMERIC {
		return createNumericCalc()
	}

	if typeCalc == COMPLEX {
		return createComplex()
	}

	return nil, fmt.Errorf("нет данного типа калькулятора")
}

func createNumericCalc() (*NumericCalculator, error) {
	calc := &NumericCalculator{
		Type: NUMERIC,
	}
	A := &models.RealNumber{}
	B := &models.RealNumber{}

	fmt.Print("Введите первое число: ")
	var a float64
	_, err := fmt.Scanln(&a)
	A.SetValue(a)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	fmt.Print("Введите второе число: ")
	var b float64
	_, err = fmt.Scanln(&b)
	B.SetValue(b)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	calc.A = *A
	calc.B = *B

	return calc, nil
}

func createComplex() (*ComplexCalculator, error) {
	fmt.Print("Введите действительную часть для первого числа: ")
	var a float64
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	fmt.Print("Введите мнимую часть для первого числа: ")
	var b float64
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	A := &models.Complex{
		ValidPart:     a,
		ImaginaryPart: b,
	}

	fmt.Print("Введите действительную часть для второг числа: ")
	var c float64
	_, err = fmt.Scanln(&c)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	fmt.Print("Введите мнимую часть для второго числа: ")
	var d float64
	_, err = fmt.Scanln(&d)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return nil, err
	}

	B := &models.Complex{
		ValidPart:     c,
		ImaginaryPart: d,
	}

	calc := &ComplexCalculator{
		A:    *A,
		B:    *B,
		Type: COMPLEX,
	}
	return calc, nil
}
