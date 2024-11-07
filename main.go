package main

import (
	"fmt"

	"example.com/note-app/note"
)

func main() {
	title, content := getNoteData()
  userNote, err := note.New(title, content)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("user note is", userNote)
}

func getUserInput(prompt string) (string) {
	var value string
	fmt.Println(prompt)
	fmt.Scanln(&value)
	return value 
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}
