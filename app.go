package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/note/structs"
)

func main() {
	title, content := getNoteData()

	userNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	userNote.Display()

	error = userNote.Save()

	if error != nil {
		fmt.Println("Saving the note failed: ", error)
		panic(error)
	}

	fmt.Println("Saving the note succeeded")
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
