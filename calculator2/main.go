package main

import (
	"fmt"
	"math"
	"strconv"
)

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

func round(num float64, decimalPlaces int) float64 {
	precision := math.Pow(10, float64(decimalPlaces))
	return math.Round(num*precision) / precision
}

func main() {
	var calc Calculator

	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	var result float64

	switch operator {
	case "+":
		result = calc.Add(a, b)
	case "-":
		result = calc.Subtract(a, b)
	case "*":
		result = calc.Multiply(a, b)
	case "/":
		result, err = calc.Divide(a, b)
	default:
		fmt.Println("Некорректный оператор")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	roundedResult := round(result, 8)
	resultStr := strconv.FormatFloat(roundedResult, 'f', -1, 64)[:9]

	fmt.Println("Результат:", resultStr)
}
