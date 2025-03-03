package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	result2 := add(1.1, 2.2)
	fmt.Println(result2)

	result3 := add("Hello, ", "World!")
	fmt.Println(result3)
}
