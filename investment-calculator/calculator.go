package main

import (
	"fmt"
	"math"
)

func outputText(name string, value *float64) {
	fmt.Printf("Enter the %s amount: ", name)
	fmt.Scan(value)
}

func main() {
	fmt.Println("Investment Calculator")
	const inflationRate = 3.5
	var interestRate, years, investmentAmount float64

	// Get input values from user
	outputText("investment", &investmentAmount)
	outputText("years", &years)
	outputText("interest rate", &interestRate)

	futureValue := investmentAmount * math.Pow((1+interestRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("Future value of investment is: %.2f\n", futureValue)
	fmt.Printf("Future real value of investment is: %.2f\n", futureRealValue)
}
