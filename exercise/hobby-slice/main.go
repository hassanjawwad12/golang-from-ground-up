package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	fmt.Println("Exercise 2")
	hobbies := [3]string{"reading", "running", "swimming"}
	fmt.Println("The values of array hobby are: ", hobbies)
	fmt.Println("The length of array hobby is: ", len(hobbies))
	fmt.Println("The capacity of array hobby is: ", cap(hobbies))
	fmt.Println("The first element of array hobby is: ", hobbies[0])

	newHobby := hobbies[1:3]
	fmt.Println("The values of array newHobby are: ", newHobby)

	mainHobbies := hobbies[:2]
	fmt.Println("The values of array mainHobbies are: ", mainHobbies)

	mainHobbies = hobbies[1:3]
	fmt.Println("The values of array mainHobbies are: ", mainHobbies)

	goals := []string{"learning", "accomplishing", "climbing"}
	fmt.Println("The values of slice goals are: ", goals)

	goals[1] = "Learn all the details!"
	goals = append(goals, "Learn all the basics!")
	fmt.Println("The values of slice goals are: ", goals)

	// we are creating a slice of products
	products := []Product{
		{Name: "laptop", Price: 2000},
		{Name: "phone", Price: 1000},
	}
	fmt.Println("The values of slice products are: ", products)

	products = append(products, Product{Name: "tablet", Price: 500})
	fmt.Println("The values of slice products are: ", products)

}
