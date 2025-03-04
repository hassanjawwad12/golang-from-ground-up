package main

import "fmt"

type transformFm func(int) int

func transformNumbers(numbers *[]int, transform transformFm) []int {
	dnumbers := []int{}
	for _, num := range *numbers {
		dnumbers = append(dnumbers, transform(num))
	}
	return dnumbers
}

// returns a function which will be further passed to the transfonumbers function
func getTransformer(num *[]int) transformFm {
	if (*num)[0] == 1 {
		return doubleInt
	} else {
		return tripleInt
	}
}

func doubleInt(num int) int {
	return num * 2
}

func tripleInt(num int) int {
	return num * 3
}

func main() {
	fmt.Println("Functions in Go")
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}
	dnumber := transformNumbers(&numbers, getTransformer(&numbers))
	fmt.Println(dnumber)

	transformerFunc := getTransformer(&moreNumbers)
	tnumber := transformNumbers(&moreNumbers, transformerFunc)
	fmt.Println(tnumber)
}
