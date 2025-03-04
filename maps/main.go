package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	fmt.Println("Maps in Go")
	fmt.Println("------------")
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	// empty map with make , we can pass only one extra value
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println()
	courseRatings.output()
	fmt.Println()

	// empty slice with make as it tells the length of the slice
	usernames2 := make([]string, 2, 5)
	usernames2[0] = "JOHN"
	usernames2 = append(usernames2, "MAX")
	usernames2 = append(usernames2, "MARK")
	usernames2 = append(usernames2, "MARK")
	for i, username := range usernames2 {
		fmt.Println(i, username)
	}
	fmt.Println("Length of usernames2:", len(usernames2))
}
