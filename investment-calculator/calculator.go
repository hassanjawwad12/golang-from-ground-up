package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")

	const inflationRate = 3.5
	interestRate := 5.5
	var years, investmentCalculator float64

	// All the initial values will be over-written
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentCalculator)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the interest rate: ")
	fmt.Scan(&interestRate)

	futureValue := investmentCalculator * math.Pow((1+interestRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	fmt.Println("Future value of investment is: ", futureValue)
	fmt.Println("Future real value of investment is: ", futureRealValue)
}
