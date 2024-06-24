package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type output interface {
	saver
	Display()
}

func saveData(data saver) error {
	error := data.Save()

	if error != nil {
		return error
	}

	return nil
}

func outputData(data output) error {
	data.Display()
	return saveData(data)
}

// func genericsAdd[T any] (value1, value2 T) T {
// 	return value1 + value2
// }

// func genericsAdd[T int | string | float64](value1, value2 T) T {
// 	return value1 + value2
// }

func printSomething(value any) {
	// typedValue, isValueAString := value.(string)
	// fmt.Println("The value", typedValue, "is of type string: ", isValueAString)

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	case todo.Todo:
		fmt.Println("Todo: ", value)
	case note.Note:
		fmt.Println("Note: ", value)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	userNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	userTodo, error := todo.New(todoText)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	error = outputData(userNote)

	if error != nil {
		fmt.Println("Saving the note failed: ", error)
		panic(error)
	}

	fmt.Println("Saving the note succeeded")

	error = outputData(userTodo)

	if error != nil {
		fmt.Println("Saving the todo failed: ", error)
		panic(error)
	}

	fmt.Println("Saving the todo succeeded")

	printSomething(1)
	printSomething(2.4)
	printSomething("Hello, world!")
	printSomething(userNote)
	printSomething(userTodo)
}

func getNoteData() (title string, content string) {
	title = getUserInput("Note Title: ")
	content = getUserInput("Note Content: ")

	return
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')

	if error != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
