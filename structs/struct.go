package main

import (
	"fmt"
	"time"
)

// define your own user type
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct method to display user data
func (u *user) displayUser() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Created At:", u.createdAt)
}

func displayUser(u *user) {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Created At:", u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// create a new user
	appuser := user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}

	// call function
	displayUser(&appuser)
	fmt.Println("")
	// call method of struct
	appuser.displayUser()
}
