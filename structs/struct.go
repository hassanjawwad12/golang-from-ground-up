package main

import (
	"fmt"

	"github.com/hassanjawwad12/golang-from-ground-up/user"
)

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

	appuser, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appuser.DisplayUser()
	appuser.ChangeName()
	appuser.DisplayUser()
}
