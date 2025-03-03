package main

import "fmt"

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println(value)
	case float64:
		fmt.Println(value)
	case string:
		fmt.Println(value)
	}

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		fmt.Println()
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		fmt.Println()
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		fmt.Println()
	}
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	// this wont be printed
	printSomething(true)
}
