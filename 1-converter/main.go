package main

import "fmt"

const (
	usdToEur = 0.8494
	usdToRub = 75.22
)

func main() {
	amount := readInput("Введите сумму: ")

	from := readCurrency("Из какой валюты (USD/EUR): ")
	to := readCurrency("В какую валюту (EUR/RUB): ")

	result := convert(amount, from, to)
	fmt.Printf("Результат конвертации (заглушка): %.2f\n", result)

	eurToRub := (1 / usdToEur) * usdToRub
	fmt.Printf("На основе констант: 1 EUR = %.2f RUB\n", eurToRub)
}

func readInput(prompt string) float64 {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func readCurrency(prompt string) string {
	var currency string
	fmt.Print(prompt)
	fmt.Scan(&currency)
	return currency
}

func convert(amount float64, from, to string) float64 {
	return 0
}
