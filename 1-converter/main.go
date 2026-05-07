package main

import "fmt"

func main() {
	const usdToEur = 0.8494
	const usdToRub = 75.22

	eurToRub := (1 / usdToEur) * usdToRub

	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
}
