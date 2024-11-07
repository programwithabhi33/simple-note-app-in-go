package note

import (
	"errors"
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
