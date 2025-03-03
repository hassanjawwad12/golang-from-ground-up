package main

import "fmt"

func editAge(age *int) {
	// does not create a new memory address because address is passed instead of value
	*age = *age + 18
}

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Deferred Age:", *agePointer)
	editAge(agePointer)

	fmt.Println("Future Age:", age)
}
