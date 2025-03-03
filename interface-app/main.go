package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hassanjawwad12/golang-from-ground-up/note"
	"github.com/hassanjawwad12/golang-from-ground-up/todo"
)

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

	todos, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	todos.Display()

	err = todos.SaveToFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Todo saved to file")

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newNote.Display()

	err = newNote.SaveToFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Note saved to file")
}
