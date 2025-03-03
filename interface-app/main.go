package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hassanjawwad12/golang-from-ground-up/note"
	"github.com/hassanjawwad12/golang-from-ground-up/todo"
)

// any struct that adheres to this interface must have a Save method
// so we can pass any struct to the function expecting a saver interface which have a save method
type saver interface {
	Save() error
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// read the user input from command line os.Stdin
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func main() {
	fmt.Println("Welcome to the interface app")
	title, content := getNoteData()
	todoText := getUserInput("Todo content:")

	// Creating and saving Todo
	todos, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	todos.Display()
	err = saveData(todos)

	if err != nil {
		return
	}
	fmt.Println("Todo saved to file")

	// Creating and saving Note
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newNote.Display()
	err = saveData(newNote)

	if err != nil {
		return
	}
	fmt.Println("Note saved to file")
}
