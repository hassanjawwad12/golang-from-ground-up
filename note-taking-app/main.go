package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hassanjawwad12/golang-from-ground-up/note"
)

func main() {
	title, content := getNoteData()

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newNote.Display()
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
