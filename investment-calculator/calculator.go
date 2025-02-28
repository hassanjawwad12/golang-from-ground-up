package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")

	investmentCalculator, years, interestRate := 1000.0, 10.0, 5.5

	futureValue := investmentCalculator * math.Pow((1+interestRate/100), years)
	fmt.Println("Future value of investment is: ", futureValue)
}
