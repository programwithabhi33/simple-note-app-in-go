package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note-app/note"
	"example.com/note-app/todo"
)

func main() {
	title, content := getNoteData()
  todoText := getUserInput("Enter your todo text: ")
  todo, err := todo.New(todoText)

  if err != nil {
    fmt.Println(err)
    return
  }

  todo.Display()

  err = todo.Save()
  if err != nil {
    fmt.Println("Saving todo failed due to some reason")
    return
  }
  fmt.Println("Todo saved successfully!")

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
