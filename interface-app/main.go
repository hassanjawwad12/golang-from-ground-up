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

// embedding the saver interface in the outputtable interface
type outputtable interface {
	saver
	Display()
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

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
	title, content := getNoteData()
	todoText := getUserInput("Todo content:")

	todos, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = outputData(newNote)
	if err != nil {
		return
	}
	fmt.Println("Note saved to file")

	err = outputData(todos)
	if err != nil {
		return
	}
	fmt.Println("Todo saved to file")
}
