package main

import "fmt"

const (
	usdToEur = 0.8494
	usdToRub = 75.22
)

func main() {
	fmt.Println("=== Конвертер валют (USD, EUR, RUB) ===")

	from := readCurrency("Введите исходную валюту (USD/EUR/RUB): ")
	amount := readAmount("Введите сумму: ")
	to := readCurrency("Введите целевую валюту (USD/EUR/RUB): ")
	result := convert(amount, from, to)

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}

func readCurrency(prompt string) string {
	var cur string
	for {
		fmt.Print(prompt)
		fmt.Scan(&cur)
		if isValidCurrency(cur) {
			return cur
		}
		fmt.Println("Ошибка: допустимые валюты - USD, EUR, RUB. Попробуйте снова.")
	}
}

func convert(amount float64, from, to string) float64 {
	var inUSD float64
	switch from {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / usdToEur
	case "RUB":
		inUSD = amount / usdToRub
	}

	switch to {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * usdToEur
	case "RUB":
		return inUSD * usdToRub
	}
	return 0
}

func readAmount(prompt string) float64 {
	var amount float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка: введите положительное число. Попробуйте снова.")
		var discard string
		fmt.Scan(&discard)
	}
}

func isValidCurrency(currency string) bool {
	return currency == "USD" || currency == "EUR" || currency == "RUB"
}
