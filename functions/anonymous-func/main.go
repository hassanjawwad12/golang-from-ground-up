package main

import "fmt"

type transformFm func(int) int

func transformNumbers(numbers *[]int, transform transformFm) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func main() {
	numbers := []int{1, 2, 3}

	// anonymous func defined as an argument to the transformNumbers function
	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)
}
