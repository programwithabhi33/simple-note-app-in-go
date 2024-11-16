package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
)

func main() {
	title, content := getNoteData()
  userNote, err := note.New(title, content)
  if err != nil {
    fmt.Println(err)
    return
  }
  userNote.Display()
  err = userNote.Save()
  if err != nil {
    fmt.Println("Saving note failed due to some reason")
    return
  }
  fmt.Println("Note saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

  reader := bufio.NewReader(os.Stdin)
  text, err := reader.ReadString('\n') 
  if err != nil {
    return ""
  }
  text = strings.TrimSuffix(text, "\n")
  text = strings.TrimSuffix(text, "\r")

	return text 
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}
