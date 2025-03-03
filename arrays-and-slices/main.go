package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println("values in array products are: ", productNames)

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("values in array prices are: ", prices)
	fmt.Println("Second index of array prices is: ", prices[2])

	// starting index is inclusive, ending index is exclusive
	// featuredPrices is a slice
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	// all values after the 1st index
	featuredPrices2 := prices[1:]
	fmt.Println("All values of array prices after 0th index : ", featuredPrices2)

	// this overwrites the value in the original array
	featuredPrices[0] = 5.99
	fmt.Println("values in array prices are: ", prices)

	// all values before the 2nd index
	highlightedPrices := featuredPrices2[:1]
	fmt.Println("All values of array featuredPrices2 before 1st index : ", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("values in array highlighter prices are: ", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = append(highlightedPrices, 5.99)
	fmt.Println("values in array highlighter prices are: ", highlightedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	discountPrices = append(highlightedPrices, discountPrices...)
	fmt.Println(discountPrices)
}
