package user

import (
	"errors"
	"fmt"
	"time"
)

// define your own user type
type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	User     //embedding
	Email    string
	Password string
}

// struct constructor (just a convention)
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
			CreatedAt: time.Now(),
		},
	}
}

func (u *User) DisplayUser() {
	fmt.Println("First Name:", u.FirstName)
	fmt.Println("Last Name:", u.LastName)
	fmt.Println("Birth Date:", u.BirthDate)
	fmt.Println("Created At:", u.CreatedAt)
}

func (u *User) ChangeName() {
	u.FirstName = "Mark"
}
