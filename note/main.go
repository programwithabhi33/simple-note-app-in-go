package note

import (
	"errors"
	"fmt"
	"time"
)

type note struct{
  title string
  content string
  createdAt time.Time
}

func New(title, content string) (note, error) {
  if title == "" || content == "" {
    return note{}, errors.New("Invalid Input")
  }
  return note{
    title, 
    content,
    time.Now(),
  }, nil
}

func (note note) Display(){
  fmt.Printf("Your note titled %v is the following content\n\n%v\n\n", note.title, note.content)
}
