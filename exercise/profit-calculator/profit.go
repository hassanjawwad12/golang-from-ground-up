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
	fmt.Println("Profit is: ", profit)
	fmt.Println("Tax is: ", tax)
	fmt.Println("Net profit is: ", netProfit)

}
