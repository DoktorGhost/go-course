package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	decimalSum, err := DecimalSum("11.5234", "0.5647")
	if err != nil {
		fmt.Println("decimalSum error: ", err)
	} else {
		fmt.Println("decimalSum: ", decimalSum)
	}

	decimalSubtract, err := DecimalSubtract("11.53452", "0.56234")
	if err != nil {
		fmt.Println("decimalSubtract error: ", err)
	} else {
		fmt.Println("decimalSubtract: ", decimalSubtract)
	}

	decimalMultiply, err := DecimalMultiply("11.53234", "0.5234")
	if err != nil {
		fmt.Println("decimalMultiply error: ", err)
	} else {
		fmt.Println("decimalMultiply: ", decimalMultiply)
	}

	decimalDivide, err := DecimalDivide("11.5345", "0.0")
	if err != nil {
		fmt.Println("decimalDivide error: ", err)
	} else {
		fmt.Println("decimalDivide: ", decimalDivide)
	}

	decimalRound, err := DecimalRound("11.545384563", 3)
	if err != nil {
		fmt.Println("decimalRound error: ", err)
	} else {
		fmt.Println("decimalRound: ", decimalRound)
	}

	decimalGreaterThan, err := DecimalGreaterThan("12.123234", "12.124234")
	if err != nil {
		fmt.Println("decimalGreaterThan error: ", err)
	} else {
		fmt.Println("decimalGreaterThan: 12.123 > 12.124 ", decimalGreaterThan)
	}

	decimalLessThan, err := DecimalLessThan("12.123", "12.124")
	if err != nil {
		fmt.Println("decimalLessThan error: ", err)
	} else {
		fmt.Println("decimalLessThan: 12.123 < 12.124 ", decimalLessThan)
	}
	decimalEqual, err := DecimalEqual("12.123", "12.124")
	if err != nil {
		fmt.Println("decimalEqual error: ", err)
	} else {
		fmt.Println("decimalEqual: 12.123 == 12.124 ", decimalEqual)
	}

}

func DecimalSum(a, b string) (string, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.Add(value2)

	return result.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.Sub(value2)
	return result.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.Mul(value2)
	return result.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	if value2.IsZero() {
		return "", fmt.Errorf("деление на ноль недопустимо")
	}
	result := value1.Div(value2)
	return result.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	value, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	result := value.Round(precision)
	return result.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.GreaterThan(value2)
	return result, nil
}

func DecimalLessThan(a, b string) (bool, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.LessThan(value2)
	return result, nil

}

func DecimalEqual(a, b string) (bool, error) {
	value1, err := decimal.NewFromString(a)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 1: %v", err)
	}

	value2, err := decimal.NewFromString(b)
	if err != nil {
		return false, fmt.Errorf("ошибка при преобразовании чилса 2: %v", err)
	}

	result := value1.Equal(value2)
	return result, nil

}
