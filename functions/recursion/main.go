package main

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120

func main() {
	result := factorial(5)
	println(result)

	result = fibonacci(7)
	println(result)
}
