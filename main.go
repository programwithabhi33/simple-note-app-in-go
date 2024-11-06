package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
    fmt.Println(err)
    return
	}
  fmt.Println(title, content)
}

func getUserInput(prompt string) (string, error) {
	var value string
	fmt.Println(prompt)
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("Note cannot to blank")
	}
	return value, nil
}
func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")
	if err != nil {
		return "", "", err
	}
	content, err := getUserInput("Note content:")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}
