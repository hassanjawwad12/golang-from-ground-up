package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("content cannot be empty")
	}
	return &Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println("Your todo has the following content: ", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, json, 0644)

	return nil
}
