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

// struct constructor (just a convention)
func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

// struct method to display user data
func (u *user) displayUser() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Created At:", u.createdAt)
}

// struct method to change user name
func (u *user) changeName() {
	u.firstName = getUserData("Please enter your first name again: ")
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

	// create a new user using constructor
	appuser := newUser(firstName, lastName, birthdate)
	appuser.displayUser()
	appuser.changeName()
	appuser.displayUser()
}
