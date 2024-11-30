package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note-app/note"
	"example.com/note-app/todo"
)
type saver interface {
  Save() error
}

/*type displayer interface {
  Display()
}*/

type outputtable interface {
  saver
  Display()
}

func saveData(data saver) error {
  err := data.Save()
  if err != nil {
    fmt.Println("Saving data failed due to some reason")
    return err
  }
  fmt.Println("Data saved successfully!")
  return nil
}

func outputData(data outputtable) error {
  data.Display()
  return saveData(data)
}

func main() {
	title, content := getNoteData()
  todoText := getUserInput("Enter your todo text: ")

  todo, err := todo.New(todoText)
  if err != nil {
    fmt.Println(err)
    return
  }
  err = outputData(todo)
  if err != nil {
    return
  }

  userNote, err := note.New(title, content)
  if err != nil {
    fmt.Println(err)
    return
  }
  outputData(userNote)
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
