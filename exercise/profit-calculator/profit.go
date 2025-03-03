package main

import (
	"errors"
	"fmt"
	"os"
)

func outputText(name string) (float64, error) {
	var value float64
	fmt.Printf("Enter the %s : ", name)
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("invalid input")
	}
	return value, nil
}

func performCalculation(revenue, expenses, taxRate float64) (profit, tax, netProfit, ratio float64) {
	profit = revenue - expenses
	tax = profit * taxRate / 100
	netProfit = profit - tax
	ratio = netProfit / profit
	return profit, tax, netProfit, ratio
}

func storeResults(profit, tax, netProfit, ratio float64) (float64, float64, float64, float64) {
	data := fmt.Sprintf("Profit is %.2f\nTax is %.2f\nNet profit is %.2f\nNet profit ratio is %.2f\n", profit, tax, netProfit, ratio)
	os.WriteFile("profit.txt", []byte(data), 0644)
	return profit, tax, netProfit, ratio
}

func main() {

	fmt.Println("Profit Calculator")

	revenue, err := outputText("Revenue")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := outputText("Expenses")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := outputText("Tax Rate")
	if err != nil {
		fmt.Println(err)
		return
	}

	profit, tax, netProfit, ratio := performCalculation(revenue, expenses, taxRate)

	fmt.Printf("Profit is %.2f\nTax is %.2f\nNet profit is %.2f\nNet profit ratio is %.2f\n", profit, tax, netProfit, ratio)
	storeResults(profit, tax, netProfit, ratio)
}
