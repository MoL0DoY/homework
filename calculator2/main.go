package main

import (
	"fmt"
	"go2/calculator2/models"
	"strings"
)

func round(memoRealNumberLength int, realNumber string, wholeNumber string) string {
	realNumberArr := []byte{}
	if memoRealNumberLength > len(realNumber) {
		memoRealNumberLength = len(realNumber)
	}
	for i := 0; i < memoRealNumberLength; i++ {
		realNumberArr = append(realNumberArr, realNumber[i])
	}
	realNumberStr := string(realNumberArr)
	return strings.Join([]string{wholeNumber, realNumberStr}, ".")
}

func main() {
	var operator string

	fmt.Print("Введите первое число: ")
	var a float64
	_, err := fmt.Scanln(&a)

	A := &models.RealNumber{}
	A.SetValue(a)
	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	fmt.Print("Введите второе число: ")
	var b float64
	B := &models.RealNumber{}
	_, err = fmt.Scanln(&b)
	B.SetValue(b)
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

	var result *models.RealNumber

	switch operator {
	case "+":
		result = A.Add(*B)
	case "-":
		result = A.Sub(*B)
	case "*":
		result = A.Mul(*B)
	case "/":
		result = A.Div(*B)
	default:
		fmt.Println("Некорректный оператор")
		return
	}

	numberStr := fmt.Sprint(result.Value())

	res1 := strings.Split(numberStr, ".")

	wholeNumber := res1[0]

	lengthWholeNumber := len(wholeNumber)
	if lengthWholeNumber > 8 {
		fmt.Errorf("Переполнение")
		return
	}

	var result1 string
	memoRealNumberLength := 8 - lengthWholeNumber
	if memoRealNumberLength == 0 {
		result1 = wholeNumber
	} else {
		var realNumber string
		if len(res1) > 1 {
			realNumber = res1[1]
			result1 = round(memoRealNumberLength, realNumber, wholeNumber)
		} else {
			result1 = wholeNumber
		}

	}

	fmt.Println("Результат:", result1)
}
