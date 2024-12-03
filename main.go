package main

import (
	"bufio"
	"example.com/note-app/note"
	"example.com/note-app/todo"
	"fmt"
	"os"
	"strings"
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

func printAnyValue(value interface{}) {
  intVal, isInt := value.(int)
  if isInt {
    fmt.Println("Integer: ", intVal)
    return
  }

  float64Val, isFloat64 := value.(float64)
  if isFloat64 {
    fmt.Println("Float: ", float64Val)
    return
  }

  stringVal, isString := value.(string)
  if isString {
    fmt.Println("String: ", stringVal)
    return
  }
	/*switch value.(type) {
	case int:
		fmt.Println("Int: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Print("String", value)
	default:
		fmt.Println("Default Case: ", value)
	}*/
}

func add[T int|float64|string](a, b T) T{
  return a + b 
}

func main() {
  result := add(4,5)
  fmt.Println("Result + 1 is: ", result + 1)

	printAnyValue(1)
	printAnyValue(1.1)
	printAnyValue("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Enter your todo text: ")

	todo, err := todo.New(todoText)
	printAnyValue(todo)
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
