package main

import "fmt"

func main() {

	fmt.Println("Profit Calculator")
	var revenue, expenses, taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	profit := revenue - expenses
	tax := profit * taxRate / 100
	netProfit := profit - tax
	ratio := netProfit / profit

	fmt.Printf("Profit is %.2f\nTax is %.2f\nNet profit is %.2f\nNet profit ratio is %.2f\n", profit, tax, netProfit, ratio)
}
