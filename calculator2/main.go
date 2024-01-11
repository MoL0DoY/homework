package main

import (
	"fmt"
	"go2/calculator2/calculators"
)

func main() {
	var operator string

	fmt.Println("Введите тип калькулятора")

	var tc string

	_, err := fmt.Scanln(&tc)

	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	calc, err := calculators.NewCalculator(tc)

	if err != nil {
		fmt.Println("Ошибка создания:", err)
		return
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	result1, err := calc.Execute(operator)

	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	fmt.Println("Результат:", result1)
}
