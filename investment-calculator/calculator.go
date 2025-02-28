package main

import (
	"fmt"
	"math"
)

// Define globally when used at more than one instance
const inflationRate = 3.5

func outputText(name string, value *float64) {
	fmt.Printf("Enter the %s : ", name)
	fmt.Scan(value)
}

func calculateFutureValue(investmentAmount, interestRate, years float64) (futureValue float64, realFutureValue float64) {
	futureValue = investmentAmount * math.Pow((1+interestRate/100), years)
	realFutureValue = futureValue / math.Pow((1+inflationRate/100), years)
	return
}

func main() {
	var interestRate, years, investmentAmount float64

	outputText("Investment", &investmentAmount)
	outputText("Years", &years)
	outputText("Interest Rate", &interestRate)

	a, b := calculateFutureValue(investmentAmount, interestRate, years)

	fmt.Printf("Future value of investment is: %.2f\n", a)
	fmt.Printf("Future real value of investment is: %.2f\n", b)
}
