package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gillgillint/structs-practice/note"
	"github.com/gillgillint/structs-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputTable interface {
	saver
	displayer
	// Display()
}

// type outputTable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content

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

func outputData(data outputTable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(int)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(int)
	if ok {
		fmt.Println("String::", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("int:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }

}
