package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type note struct{
  Title string `json:"title"`
  Content string `json:"content"`
  CreatedAt time.Time `json:"created_at"` 
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

func (note note) Display() {
  fmt.Printf("Your note titled %v is the following content\n\n%v\n\n", note.Title, note.Content)
}

func (note note) Save() error {
  fileName := strings.ReplaceAll(note.Title, " ", "_")
  fileName = strings.ToLower(fileName) + ".json"
  
  json, err := json.Marshal(note)
  if err != nil {
    return err
  }
  return os.WriteFile(fileName, json, 0644)
}
