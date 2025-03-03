package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content cannot be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}
